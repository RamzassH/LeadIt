package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/internal/domain/models"
	"github.com/RamzassH/LeadIt/authService/internal/lib/jwt"
	"github.com/RamzassH/LeadIt/authService/internal/lib/verification"
	"github.com/RamzassH/LeadIt/authService/internal/storage"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrTokenNotFound      = errors.New("token not found")
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneRegex = regexp.MustCompile(`^\+?[0-9\s\-]+$`)
)

func isEmail(identifier string) bool {
	return emailRegex.MatchString(identifier)
}

func isPhoneNumber(identifier string) bool {
	return phoneRegex.MatchString(identifier)
}

type Auth struct {
	logger          zerolog.Logger
	userSaver       UserSaver
	userProvider    UserProvider
	tokenSaver      TokenSaver
	redisStorage    Redis
	kafka           *kafka.Producer
	tokenTTL        time.Duration
	refreshTokenTTL time.Duration
}

type ConfirmMassage struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Code   string `json:"code"`
	Time   string `json:"timestamp"`
}

type TokenSaver interface {
	SaveRefreshToken(ctx context.Context, token models.RefreshToken) error
	GetRefreshToken(ctx context.Context, token string) (models.RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, token string) error
}

type UserSaver interface {
	SaveUser(ctx context.Context,
		user models.User,
	) (uid int64, err error)
}

type UserProvider interface {
	UserByEmail(ctx context.Context, email string) (user models.User, err error)
	UserById(ctx context.Context, id int64) (user models.User, err error)
	IsAdmin(ctx context.Context, uid int64) (isAdmin bool, err error)
	UpdateUser(ctx context.Context, user models.UpdateUserPayload) error
	VerifyUser(ctx context.Context, userID int64) error
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}

const VerificationCodeTTL = 15 * time.Minute

func New(
	log zerolog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	tokenSaver TokenSaver,
	redisStorage Redis,
	kafka *kafka.Producer,
	tokenTTL time.Duration,
	refreshTokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:       userSaver,
		userProvider:    userProvider,
		logger:          log,
		tokenSaver:      tokenSaver,
		redisStorage:    redisStorage,
		kafka:           kafka,
		tokenTTL:        tokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string) (token string, refreshToken string, err error) {
	const op = "auth.Login"

	logger := a.logger.With().Str("operation", op).Logger()

	a.logger.Info().Str("operation", op).Msg("login user")

	user, err := a.userProvider.UserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			logger.Error().Msg("user not found")

			return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		} else {
			logger.Error().Err(err).Msg("failed to fetch user")

			return "", "", fmt.Errorf("%s: %w", op, err)
		}
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		logger.Error().Err(err).Msg("invalid credentials")

		return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	token, err = jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate token")

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err = generateSecureToken()
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	rt := models.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(a.refreshTokenTTL),
		Revoked:   false,
		CreatedAt: time.Now(),
	}

	if err := a.tokenSaver.SaveRefreshToken(ctx, rt); err != nil {
		logger.Error().Err(err).Msg("failed to save refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return token, refreshToken, nil
}

func (a *Auth) Logout(ctx context.Context, refreshToken string) error {
	const op = "auth.Logout"

	logger := a.logger.With().Str("operation", op).Logger()

	err := a.tokenSaver.RevokeRefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, storage.ErrTokenNotFound) {
			return fmt.Errorf("%s: %w", op, storage.ErrTokenNotFound)
		}
		logger.Error().Err(err).Msg("failed to revoke refresh token")
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, user models.RegisterUserPayload) (userId int64, err error) {
	const op = "auth.Register"

	logger := a.logger.With().Str("operation", op).Logger()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error().Err(err).Str(op, err.Error())
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().Str("username", user.Name).Str("surname", user.Surname).Str("email", user.Email).Msg("register user")

	var id int64
	var saveErr error

	userEntity := models.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		PassHash: hashPass,
	}
	id, saveErr = a.userSaver.SaveUser(ctx, userEntity)
	if saveErr != nil {
		logger.Error().Err(saveErr).Str("operation", op).Msg("failed to save user")
		return 0, fmt.Errorf("%s: %w", op, saveErr)
	}

	code, err := verification.GenerateVerificationCode()
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate verification code")
	}

	userVerificationKey := fmt.Sprintf("auth:verefication:%s", code)
	logger.Info().Str("key stored set", userVerificationKey).Msg("key stored")

	redisErr := a.redisStorage.Set(ctx, userVerificationKey, userEntity.Email, VerificationCodeTTL)

	if redisErr != nil {
		logger.Error().Err(err).Msg("failed to set user verification code")
		return 0, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	storedKey, err := a.redisStorage.Get(ctx, userVerificationKey)

	logger.Info().Str("get key", storedKey).Msg("getKey")

	ConfirmMsg := ConfirmMassage{
		UserID: id,
		Email:  user.Email,
		Code:   code,
		Time:   time.Now().UTC().Format(time.RFC3339),
	}

	msgBytes, err := json.Marshal(ConfirmMsg)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal confirm message")
	} else {
		if err := a.kafka.Send(ctx, []byte("notification"), msgBytes); err != nil {
			logger.Error().Err(err).Msg("failed to send confirm message")
		} else {
			logger.Info().Msg("confirm message sent to Kafka")
		}
	}

	return id, nil
}

func (a *Auth) UpdateUser(ctx context.Context, updatePayload models.UpdateUserPayload) error {
	const op = "Auth.UpdateUser"

	var hashedPassword []byte
	if updatePayload.Password != "" {
		var err error
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(updatePayload.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("%s: password hashing failed: %w", op, err)
		}
	}

	user := updatePayload
	user.Password = string(hashedPassword)

	if err := a.userProvider.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *Auth) VerifyCode(ctx context.Context, verifyPayload models.VerifyUserPayload) (token string, refreshToken string, err error) {
	const op = "auth.VerifyCode"

	logger := a.logger.With().Str("operation", op).Logger()

	key := fmt.Sprintf("auth:verefication:%s", verifyPayload.Code)
	storedEmail, err := a.redisStorage.Get(ctx, key)
	logger.Info().Str("storedKey", key).Msg("Проверка хранения кода в Redis")
	logger.Info().Str("stored email", storedEmail).Msg(storedEmail)
	if err != nil {
		if errors.Is(err, storage.ErrTokenNotFound) {
			return "", "", fmt.Errorf("%s: %w", op, storage.ErrTokenNotFound)
		}
	}
	if storedEmail == "" {
		logger.Error().Msg("invalid code")
		return "", "", status.Error(codes.InvalidArgument, "invalid code")
	}

	_ = a.redisStorage.Del(ctx, key)
	logger.Info().Msg("user successfully verified")

	user, err := a.userProvider.UserByEmail(ctx, storedEmail)
	logger.Info().Interface("user", user).Msg("user found")
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return "", "", fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
	}

	token, err = jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate token")

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err = generateSecureToken()
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	rt := models.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(a.refreshTokenTTL),
		Revoked:   false,
		CreatedAt: time.Now(),
	}
	logger.Debug().Str("email", storedEmail).Str("Email", user.Email)
	logger.Debug().Str("refreshToken", refreshToken).Msg("refresh token")

	if err := a.tokenSaver.SaveRefreshToken(ctx, rt); err != nil {
		logger.Error().Err(err).Msg("failed to save refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return token, refreshToken, nil
}

func (a *Auth) IsAdmin(ctx context.Context, uid int64) (bool, error) {
	const op = "auth.IsAdmin"

	logger := a.logger.With().Str("operation", op).Logger()

	isAdmin, err := a.userProvider.IsAdmin(ctx, uid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			logger.Error().Err(err).Msg("user not found")

			return false, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

func (a *Auth) RefreshToken(ctx context.Context, refreshToken string) (token string, newRefreshToken string, err error) {
	const op = "auth.RefreshToken"

	logger := a.logger.With().Str("operation", op).Logger()

	rt, err := a.tokenSaver.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, storage.ErrTokenNotFound) {
			return "", "", fmt.Errorf("%s: %w", op, storage.ErrTokenNotFound)
		}
		logger.Error().Err(err).Msg("failed to get refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if rt.Revoked || time.Now().After(rt.ExpiresAt) {
		return "", "", fmt.Errorf("%s: invalid or expired refresh token", op)
	}

	user, err := a.userProvider.UserById(ctx, rt.UserID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get user by ID")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	token, err = jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	newRefreshToken, err = generateSecureToken()
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	newRT := models.RefreshToken{
		Token:     newRefreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(a.refreshTokenTTL),
		Revoked:   false,
		CreatedAt: time.Now(),
	}
	if err := a.tokenSaver.SaveRefreshToken(ctx, newRT); err != nil {
		logger.Error().Err(err).Msg("failed to save new refresh token")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if err := a.userProvider.VerifyUser(ctx, user.ID); err != nil {
		logger.Error().Err(err).Msg("failed to verify user")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return token, newRefreshToken, nil
}

func generateSecureToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
