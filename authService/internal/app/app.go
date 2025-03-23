package app

import (
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/authService/internal/app/grpc"
	"github.com/RamzassH/LeadIt/authService/internal/services/auth"
	"github.com/RamzassH/LeadIt/authService/internal/storage/postgreSQL"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/RamzassH/LeadIt/libs/redis"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

func New(
	logger zerolog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	validate *validator.Validate,
	db *sql.DB,
	redisClient *redis.Client,
	kafka *kafka.Producer) (*App, error) {

	storage, err := postgreSQL.New(db)
	rStorage, err := redisStorage.New(redisClient, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %w", err)
	}

	authService := auth.New(
		logger,
		storage,
		storage,
		storage,
		rStorage,
		kafka,
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
