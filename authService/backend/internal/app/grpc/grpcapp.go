package grpcapp

import (
	"fmt"
	authgrpc "github.com/RamzassH/LeadIt/authService/backend/internal/grpc/auth"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	validator  *validator.Validate
	port       int
}

func New(
	log *slog.Logger,
	port int,
	validate *validator.Validate,
	authService authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.RegisterGRPCServer(gRPCServer, validate, authService)

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
	const operation = "grpcs.Start"

	logger := application.logger.With(slog.String("operation", operation))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", application.port))
	if err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Info("Starting gRPC server", slog.String("address", l.Addr().String()))

	if err := application.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	return nil
}

func (application *App) Stop() {
	const operation = "grpcs.Stop"
	logger := application.logger.With(slog.String("operation", operation))

	logger.Info("stopping gRPC server", slog.Int("port", application.port))

	application.gRPCServer.GracefulStop()
}
