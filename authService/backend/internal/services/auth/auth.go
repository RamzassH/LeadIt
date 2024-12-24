package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/backend/internal/domain/models"
	"github.com/RamzassH/LeadIt/authService/backend/internal/lib/jwt"
	"github.com/RamzassH/LeadIt/authService/backend/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"regexp"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
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
	logger          *slog.Logger
	userSaver       UserSaver
	userProvider    UserProvider
	appProvider     AppProvider
	tokenSaver      TokenSaver
	tokenTTL        time.Duration
	refreshTokenTTL time.Duration
}

type TokenSaver interface {
	SaveRefreshToken(ctx context.Context, token models.RefreshToken) error
	GetRefreshToken(ctx context.Context, token string) (models.RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, token string) error
}

type UserSaver interface {
	SaveUser(ctx context.Context,
		name string,
		surname string,
		email string,
		passwordHash []byte,
		isAdmin bool,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (user models.User, err error)
	UserById(ctx context.Context, id int64) (user models.User, err error)
	IsAdmin(ctx context.Context, uid int64) (isAdmin bool, err error)
}

type AppProvider interface {
	App(ctx context.Context, appID int64) (app models.App, err error)
}

func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenSaver TokenSaver,
	tokenTTL time.Duration,
	refreshTokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:       userSaver,
		userProvider:    userProvider,
		logger:          log,
		appProvider:     appProvider,
		tokenSaver:      tokenSaver,
		tokenTTL:        tokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, email string, password string, appID int64) (token string, refreshToken string, err error) {
	const op = "auth.Login"

	logger := a.logger.With(slog.String("operation", op))

	logger.Info("login user")

	user, err := a.userProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.logger.Warn("user not found")

			return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		} else {
			a.logger.Error("failed to fetch user", err)

			return "", "", fmt.Errorf("%s: %w", op, err)
		}
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.logger.Info("invalid credentials", err)

		return "", "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	app, err := a.appProvider.App(ctx, appID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	token, err = jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		a.logger.Error("failed to generate token", err)

		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, err = generateSecureToken()
	if err != nil {
		a.logger.Error("failed to generate refresh token", err)
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
		a.logger.Error("failed to save refresh token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return token, refreshToken, nil
}

func (a *Auth) Logout(ctx context.Context, refreshToken string) error {
	const op = "auth.Logout"

	logger := a.logger.With(slog.String("operation", op))

	err := a.tokenSaver.RevokeRefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, storage.ErrTokenNotFound) {
			return fmt.Errorf("%s: %w", op, storage.ErrTokenNotFound)
		}
		logger.Error("failed to revoke refresh token", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, name string, surname string, email string, password string, isAdmin bool) (userID int64, err error) {
	const op = "auth.Register"

	logger := a.logger.With(
		slog.String("operation", op),
	)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(op, err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var id int64
	var saveErr error

	logger.Info("register user",
		slog.String("name", name),
		slog.String("surname", surname),
		slog.String("email", email),
		slog.Bool("isAdmin", isAdmin),
	)
	id, saveErr = a.userSaver.SaveUser(ctx, name, surname, email, hashPass, isAdmin)
	if saveErr != nil {
		if errors.Is(saveErr, storage.ErrUserExists) {
			logger.Warn("user already exists")
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		logger.Error(op, saveErr)
		logger.Info("hui")
		return 0, fmt.Errorf("%w: %s", saveErr, op)
	}

	return id, nil
}

func (a *Auth) IsAdmin(ctx context.Context, uid int64) (bool, error) {
	const op = "auth.IsAdmin"

	logger := a.logger.With(slog.String("operation", op))

	logger.Info("check user is admin")

	isAdmin, err := a.userProvider.IsAdmin(ctx, uid)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			logger.Warn("user not found")

			return false, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("check user is admin", slog.Bool("isAdmin", isAdmin))

	return isAdmin, nil
}

func (a *Auth) RefreshToken(ctx context.Context, refreshToken string) (token string, newRefreshToken string, err error) {
	const op = "auth.RefreshToken"

	logger := a.logger.With(slog.String("operation", op))

	rt, err := a.tokenSaver.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, storage.ErrTokenNotFound) {
			return "", "", fmt.Errorf("%s: %w", op, storage.ErrTokenNotFound)
		}
		logger.Error("failed to get refresh token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if rt.Revoked || time.Now().After(rt.ExpiresAt) {
		return "", "", fmt.Errorf("%s: invalid or expired refresh token", op)
	}

	user, err := a.userProvider.UserById(ctx, rt.UserID)
	if err != nil {
		logger.Error("failed to get user by ID", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	//TODO заглушка, разобраться с id приложений
	app, err := a.appProvider.App(ctx, 1)
	if err != nil {
		logger.Error("failed to get app", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	token, err = jwt.NewToken(user, app, a.tokenTTL)
	if err != nil {
		logger.Error("failed to generate token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	newRefreshToken, err = generateSecureToken()
	if err != nil {
		logger.Error("failed to generate refresh token", err)
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
		logger.Error("failed to save new refresh token", err)
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
