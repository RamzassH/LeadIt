package grpcapp

import (
	"fmt"
	authgrpc "github.com/RamzassH/LeadIt/authService/internal/grpc/auth"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	logger     zerolog.Logger
	gRPCServer *grpc.Server
	validator  *validator.Validate
	port       int
}

func New(
	log zerolog.Logger,
	port int,
	validate *validator.Validate,
	authService authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.RegisterGRPCServer(gRPCServer, validate, log, authService)

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

	logger := application.logger.Info().Str("operation", operation)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", application.port))
	if err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Str("address", l.Addr().String()).Msg("Starting gRPC server")

	if err := application.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	return nil
}

func (application *App) Stop() {
	const operation = "grpcs.Stop"
	logger := application.logger.Info().Str("operation", operation)

	logger.Int("port", application.port).Msg("stopping gRPC server")

	application.gRPCServer.GracefulStop()
}
