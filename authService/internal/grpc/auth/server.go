package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/internal/services/auth"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string) (token string, refreshToken string, err error)

	VerifyCode(
		ctx context.Context,
		email string,
		code string) (token string, refreshToken string, err error)

	RegisterNewUser(
		ctx context.Context,
		name string,
		surname string,
		email string,
		password string) (userID int64, err error)

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

type RegisterValidation struct {
	Name     string `validate:"required,min=1,max=100"`
	Surname  string `validate:"required,min=1,max=100"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type LoginRequestValidation struct {
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type VerifyValidation struct {
	Email string `json:"email,omitempty" validate:"required"`
	Code  string `json:"code,omitempty" validate:"required"`
}

type IsAdminValidation struct {
	userId int64 `validate:"required,min=1,max=100"`
}

func (s *ServerAPI) ValidateStruct(data interface{}) error {
	if err := s.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (s *ServerAPI) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	registerReq := RegisterValidation{
		Name:     req.GetName(),
		Surname:  req.GetSurname(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := s.ValidateStruct(registerReq); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.RegisterNewUser(ctx, req.GetName(), req.GetSurname(), req.GetEmail(), req.GetPassword())

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

func (s *ServerAPI) VerifyCode(ctx context.Context, req *authv1.VerifyRequest) (*authv1.VerifyResponse, error) {
	verifyReq := VerifyValidation{
		Email: req.GetEmail(),
		Code:  req.GetCode(),
	}

	if err := s.ValidateStruct(verifyReq); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, refreshToken, err := s.auth.VerifyCode(ctx, req.GetEmail(), req.GetCode())
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}

	cookie := fmt.Sprintf("access_token=%s; HttpOnly; Secure; Path=/", token)

	metaData := metadata.Pairs("Cookie", cookie)

	if err := grpc.SetHeader(ctx, metaData); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	return &authv1.VerifyResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *ServerAPI) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	loginReq := LoginRequestValidation{
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

	cookie := fmt.Sprintf("access_token=%s; HttpOnly; Secure; Path=/", token)

	metaData := metadata.Pairs("Cookie", cookie)

	if err := grpc.SetHeader(ctx, metaData); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &authv1.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *ServerAPI) IsAdmin(ctx context.Context, req *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	isAdminReq := IsAdminValidation{
		userId: req.GetUserId(),
	}

	if err := s.ValidateStruct(isAdminReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate isAdmin: %w", err)
	}

	isAdmin, err := s.auth.IsAdmin(ctx, isAdminReq.userId)
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
	panic("implement me")
}

func (s *ServerAPI) ResetPassword(ctx context.Context, req *authv1.ResetPasswordRequest) (*authv1.ResetPasswordResponse, error) {
	panic("implement me")
}
