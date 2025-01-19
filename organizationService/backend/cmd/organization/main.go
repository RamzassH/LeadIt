package main

import (
	"backend/internal/config"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()

	logger := setupLogger(cfg.Env)

	logger.Info("starting organization service", slog.String("env", cfg.Env))

	validate := validator.New()

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

	//TODO инициализация приложения
	//TODO старт сервера
	//TODO save stop
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
