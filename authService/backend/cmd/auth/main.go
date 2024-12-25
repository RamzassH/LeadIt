package main

import (
	"database/sql"
	"github.com/RamzassH/LeadIt/authService/backend/internal/app"
	"github.com/RamzassH/LeadIt/authService/backend/internal/config"
	"github.com/go-playground/validator/v10"
	"log/slog"
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

	logger.Info("Starting auth service", slog.String("env", cfg.Env))

	validate := validator.New()

	logger.Info("connection string:", cfg.PostgresDSN)
	db, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		logger.Error("Failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Failed to close database connection", slog.Any("error", err))
		}
	}()
	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database", slog.Any("error", err))
		os.Exit(1)
	}
	logger.Info("Successfully connected to the database")

	application, err := app.New(
		logger,
		cfg.GRPC.Port,
		cfg.TokenTTL,
		cfg.RefreshTokenTTL,
		validate,
		db,
	)
	if err != nil {
		logger.Error("Failed to initialize application", slog.Any("error", err))
		os.Exit(1)
	}

	go application.GRPCServer.MustStart()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	logger.Info("Shutting down...", slog.String("signal", sign.String()))

	application.GRPCServer.Stop()

	logger.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
