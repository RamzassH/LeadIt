package app

import (
	server "github.com/RamzassH/LeadIt/notificationService/internal/server/notification"
	notification "github.com/RamzassH/LeadIt/notificationService/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"kafka"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	notificationServer *server.ServerAPI
	logger             zerolog.Logger
}

func New(
	log zerolog.Logger,
	kafkaConsumer *kafka.Consumer,
	validate *validator.Validate,
	smtp notification.SMTP) (*App, error) {

	notificationService := notification.New(log, smtp)

	nServer := server.New(notificationService, log, kafkaConsumer, validate)

	return &App{
		notificationServer: nServer,
		logger:             log,
	}, nil
}

func (a *App) Start() {
	a.logger.Info().Msg("Starting application...")
	a.notificationServer.StartKafkaConsumer()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	a.logger.Info().Str("signal", sig.String()).Msg("Shutting down...")
}
