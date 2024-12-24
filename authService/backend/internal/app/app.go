package app

import (
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/backend/internal/app/grpc"
	"github.com/RamzassH/LeadIt/authService/backend/internal/services/auth"
	"github.com/RamzassH/LeadIt/authService/backend/internal/storage/postgreSQL"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

func New(
	logger *slog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	validate *validator.Validate,
	db *sql.DB) (*App, error) {

	storage, err := postgreSQL.New(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %w", err)
	}

	authService := auth.New(
		logger,
		storage, // UserSaver
		storage, // UserProvider
		storage, // AppProvider
		storage, // TokenSaver
		tokenTTL,
		refreshTokenTTL,
	)

	grpcApp := grpcapp.New(
		logger,
		grpcPort,
		validate,
		authService,
	)

	return &App{
		GRPCServer: grpcApp,
		validator:  validate,
	}, nil
}
