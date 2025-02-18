package grpcapp

import (
	"fmt"
	notificationgrpc "github.com/RamzassH/LeadIt/notificationService/internal/grpc/notification"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	logger     *zerolog.Logger
	gRPCServer *grpc.Server
	validator  *validator.Validate
	port       int
}

func New(
	log *zerolog.Logger,
	port int,
	validate *validator.Validate,
	notificationService notificationgrpc.Notification) *App {
	gRPCServer := grpc.NewServer()

	notificationgrpc.RegisterGRPCServer(gRPCServer, validate, notificationService)

	return &App{
		logger:     log,
		gRPCServer: gRPCServer,
		port:       port,
		validator:  validate,
	}
}

func (application *App) MustStart() {
	if err := application.Start(); err != nil {
		panic(err)
	}
}

func (application *App) Start() error {
	const operation = "grpc.Start"

	logger := application.logger.With().Str("operation", operation).Logger()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", application.port))
	if err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Info().
		Str("address", l.Addr().String()).
		Msg("Starting gRPC server")

	if err := application.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	return nil
}

func (application *App) Stop() {
	const operation = "grpcs.Stop"
	logger := application.logger.With().Str("operation", operation).Logger()

	logger.Info().Int("port", application.port).Msg("stopping gRPC server")

	application.gRPCServer.GracefulStop()
}
