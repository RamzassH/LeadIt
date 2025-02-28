package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/internal/domain/models"
	"github.com/RamzassH/LeadIt/authService/internal/services/auth"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string) (token string, refreshToken string, err error)

	VerifyCode(
		ctx context.Context,
		payload models.VerifyUserPayload) (token string, refreshToken string, err error)

	RegisterNewUser(
		ctx context.Context,
		user models.RegisterUserPayload) (userID int64, err error)

	UpdateUser(
		ctx context.Context,
		updatePayload models.UpdateUserPayload) (err error)

	IsAdmin(
		ctx context.Context,
		userID int64) (bool, error)

	RefreshToken(
		ctx context.Context,
		refreshToken string) (token string, newRefreshToken string, err error)

	Logout(
		ctx context.Context,
		refreshToken string) error
}

type ServerAPI struct {
	authv1.UnimplementedAuthServer
	auth     Auth
	logger   zerolog.Logger
	validate *validator.Validate
}

func RegisterGRPCServer(grpcServer *grpc.Server, validate *validator.Validate, logger zerolog.Logger, authService Auth) {
	authv1.RegisterAuthServer(grpcServer, &ServerAPI{
		validate: validate,
		auth:     authService,
		logger:   logger,
	})
}

func (s *ServerAPI) ValidateStruct(data interface{}) error {
	if err := s.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (s *ServerAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	registerReq := models.RegisterUserPayload{
		Name:     req.GetName(),
		Surname:  req.GetSurname(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := s.ValidateStruct(registerReq); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.RegisterNewUser(ctx, registerReq)

	if err != nil {
		if errors.Is(err, auth.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (s *ServerAPI) Verify(ctx context.Context, req *authv1.VerifyRequest) (*authv1.VerifyResponse, error) {
	verifyReq := models.VerifyUserPayload{
		Email: req.GetEmail(),
		Code:  req.GetCode(),
	}

	token, refreshToken, err := s.auth.VerifyCode(ctx, verifyReq)
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := setCookieHeader(ctx, token); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.VerifyResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *ServerAPI) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	loginReq := models.LoginUserPayload{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := s.ValidateStruct(loginReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate login: %w", err)
	}

	token, refreshToken, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())

	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid credentials")
		}
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := setCookieHeader(ctx, token); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &authv1.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *ServerAPI) IsAdmin(ctx context.Context, req *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	isAdminReq := models.IsAdminPayload{
		UserId: req.GetUserId(),
	}

	if err := s.ValidateStruct(isAdminReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate isAdmin: %w", err)
	}

	isAdmin, err := s.auth.IsAdmin(ctx, isAdminReq.UserId)
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "validate isAdmin: %w", err)
	}

	return &authv1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func setCookieHeader(ctx context.Context, token string) error {
	cookie := fmt.Sprintf("access_token=%s; HttpOnly; Secure; Path=/", token)

	metaData := metadata.Pairs("Set-Cookie", cookie)
	if err := grpc.SetHeader(ctx, metaData); err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return nil
}

func (s *ServerAPI) RefreshToken(ctx context.Context, req *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	refreshToken := req.GetRefreshToken()
	if refreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh_token is required")
	}

	token, newRefreshToken, err := s.auth.RefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, auth.ErrTokenNotFound) {
			return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authv1.RefreshTokenResponse{
		Token:        token,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *ServerAPI) Logout(ctx context.Context, req *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	refreshToken := req.GetRefreshToken()
	if refreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "refresh_token is required")
	}

	err := s.auth.Logout(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, auth.ErrTokenNotFound) {
			return nil, status.Error(codes.NotFound, "refresh token not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authv1.LogoutResponse{}, nil
}

func (s *ServerAPI) UpdateUser(ctx context.Context, req *authv1.UpdateUserRequest) (*authv1.UpdateUserResponse, error) {
	birthDate, err := time.Parse(time.RFC3339, req.GetBirth())

	userId, err := strconv.ParseInt(req.GetUserId(), 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	updatePayload := models.UpdateUserPayload{
		ID:         userId,
		Name:       req.GetName(),
		Surname:    req.GetSurname(),
		MiddleName: req.GetMiddleName(),
		AboutMe:    req.GetAboutMe(),
		Messengers: req.GetMessengers(),
		Email:      req.GetEmail(),
		Password:   req.GetPassword(),
		BirthDate:  birthDate,
	}

	if err := s.ValidateStruct(updatePayload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate update: %w", err)
	}

	updateErr := s.auth.UpdateUser(ctx, updatePayload)
	if updateErr != nil {
		if errors.Is(updateErr, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, updateErr.Error())
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authv1.UpdateUserResponse{Success: true}, nil
}

func (s *ServerAPI) ResetPassword(ctx context.Context, req *authv1.ResetPasswordRequest) (*authv1.ResetPasswordResponse, error) {
	panic("implement me")
}
