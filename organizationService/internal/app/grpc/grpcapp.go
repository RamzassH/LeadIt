package grpcapp

import (
	"fmt"
	organizationgrpc "github.com/RamzassH/LeadIt/organizationService/internal/grpc/organization"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
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
	validator *validator.Validate,
	organizationService organizationgrpc.Service) *App {
	gRPCServer := grpc.NewServer()

	organizationgrpc.RegisterGRPCServer(gRPCServer, validator, log, organizationService)

	return &App{
		logger:     log,
		gRPCServer: gRPCServer,
		port:       port,
		validator:  validator,
	}
}

func (application *App) MustStart() {
	if err := application.Start(); err != nil {
		panic(err)
	}
}

func (application *App) Start() error {
	const operation = "grpc.Start"

	logger := application.logger.Info().Str("operation", operation)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", application.port))
	if err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Str("address", l.Addr().String()).Msg("gRPC server starting")

	if err := application.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Str("port", fmt.Sprintf(":%d", application.port)).Msg("gRPC server started")

	return nil
}

func (application *App) Stop() {
	const operation = "grpc.Stop"
	logger := application.logger.Info().Str("operation", operation)
	logger.Int("port", application.port).Msg("gRPC server stopping")

	application.gRPCServer.GracefulStop()

	logger.Str("port", fmt.Sprintf(":%d", application.port)).Msg("gRPC server stopped")
}
