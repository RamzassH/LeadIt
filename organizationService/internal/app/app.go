package app

import (
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/app/grpc"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"kafka"
	"redisStorage"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

func New(logger zerolog.Logger,
	grpcPort int,
	validate *validator.Validate,
	db *sql.DB,
	redisClient *redis.Client,
	kafka *kafka.Producer) (*App, error) {

	storage, err := postgreSQL.New(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize postgreSQL storage: %w", err)
	}

	rStorage, redisErr := redisStorage.New(redisClient, logger)

	if redisErr != nil {
		return nil, fmt.Errorf("failed to initialize redis storage: %w", err)
	}

	organizationService := organization.New(
		logger,
		storage,
		storage,
		storage,
		rStorage,
		kafka,
	)

	grpcApp := grpcapp.New(
		logger,
		grpcPort,
		validate,
		organizationService)

	return &App{
		GRPCServer: grpcApp,
		validator:  validate,
	}, nil
}
