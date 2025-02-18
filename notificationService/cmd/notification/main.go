package main

import (
	"github.com/RamzassH/LeadIt/notificationService/internal/app"
	"github.com/RamzassH/LeadIt/notificationService/internal/config"
	notification "github.com/RamzassH/LeadIt/notificationService/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	logger.Info().Str("env", cfg.Env).Msg("Starting notification service")

	validate := validator.New()

	smtpConfig := notification.SMTP{
		Host:     cfg.SMTP.Host,
		Port:     cfg.SMTP.Port,
		User:     cfg.SMTP.User,
		Password: cfg.SMTP.Password,
	}

	application, _ := app.New(logger, cfg.GRPC.Port, validate, smtpConfig.Host, smtpConfig.Port, smtpConfig.User, smtpConfig.Password)

	go application.GRPCServer.MustStart()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	logger.Info().Str("signal", sign.String()).Msg("Shutting down...")

	application.GRPCServer.Stop()

	logger.Info().Msg("application stopped")
}

func setupLogger(env string) *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var logger zerolog.Logger

	if env == envLocal {
		logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		logger = log.Output(os.Stdout)
	}

	return &logger
}
