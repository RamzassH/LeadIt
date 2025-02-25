package main

import (
	notification "github.com/RamzassH/LeadIt/notificationService/internal/service"
	"github.com/go-playground/validator/v10"
	"os"
	"os/signal"
	"syscall"

	"github.com/RamzassH/LeadIt/notificationService/internal/app"
	"github.com/RamzassH/LeadIt/notificationService/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"kafka"
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

	brokers := []string{"kafka:9092"}
	topic := "notification"
	groupID := "notification-consumer-group"

	consumer := kafka.NewConsumer(brokers, topic, groupID, logger)
	defer consumer.Close()

	application, err := app.New(logger, consumer, validate, smtpConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize application")
	}

	go application.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	logger.Info().Str("signal", sig.String()).Msg("Shutting down...")
}

func setupLogger(env string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if env == "local" {
		return log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return log.Output(os.Stdout)
}
