package main

import (
	"database/sql"
	"github.com/RamzassH/LeadIt/organizationService/internal/app"
	"github.com/RamzassH/LeadIt/organizationService/internal/config"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"kafka"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()

	logger := setupLogger(cfg.Env)

	logger.Info().Str("env", cfg.Env).Msg("Starting organization service")

	validate := validator.New()

	db, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Fatal().Err(err).Msg("Failed to close database connection")
		}
	}()

	if err := db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping database")
		os.Exit(1)
	}
	logger.Info().
		Str("env", cfg.Env).
		Interface("cfg", cfg).
		Int("port", cfg.GRPC.Port).
		Msg("Successfully connected to the database")

	opt, err := redis.ParseURL(cfg.RedisConnectionString)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to parse redis connection string")
	}

	redisClient := redis.NewClient(opt)

	brokers := []string{"kafka:9092"}
	topic := "notification"
	producer := kafka.NewProducer(brokers, topic, logger)

	application, err := app.New(
		logger,
		cfg.GRPC.Port,
		validate,
		db,
		redisClient,
		producer)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create application")
		os.Exit(1)
	}

	go application.GRPCServer.MustStart()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	logger.Info().Str("signal", sign.String()).Msg("Shutting down...")

	application.GRPCServer.Stop()

	logger.Info().Msg("application stopped")
}

func setupLogger(env string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if env == envLocal {
		return log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return log.Output(os.Stdout)
}
