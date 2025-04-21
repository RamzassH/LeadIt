package grpcapp

import (
	"fmt"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	"github.com/RamzassH/LeadIt/organizationService/internal/config"
	"github.com/RamzassH/LeadIt/organizationService/internal/grpc/interceptors"
	organizationgrpc "github.com/RamzassH/LeadIt/organizationService/internal/grpc/organization"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"time"
)

type App struct {
	logger     zerolog.Logger
	gRPCServer *grpc.Server
	validator  *validator.Validate
	config     *config.Config
}

func New(
	log zerolog.Logger,
	cfg *config.Config,
	validator *validator.Validate,
	organizationService organizationgrpc.Service) *App {
	authClient, err := newAuthClient()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create auth client")
		return nil
	}

	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.JwtUnaryServerInterceptor(cfg.TokenSecret)))

	organizationgrpc.RegisterGRPCServer(gRPCServer, validator, log, organizationService, authClient)

	return &App{
		logger:     log,
		gRPCServer: gRPCServer,
		config:     cfg,
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

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", application.config.GRPC.Port))
	if err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Str("address", l.Addr().String()).Msg("gRPC server starting")

	if err := application.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC server: %s %w", operation, err)
	}

	logger.Str("port", fmt.Sprintf(":%d", application.config.GRPC.Port)).Msg("gRPC server started")

	return nil
}

func (application *App) Stop() {
	const operation = "grpc.Stop"
	logger := application.logger.Info().Str("operation", operation)
	logger.Int("port", application.config.GRPC.Port).Msg("gRPC server stopping")

	application.gRPCServer.GracefulStop()

	logger.Str("port", fmt.Sprintf(":%d", application.config.GRPC.Port)).Msg("gRPC server stopped")
}

func newAuthClient() (authv1.AuthClient, error) {
	conn, err := grpc.Dial(
		"auth-service:57442",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(3*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return authv1.NewAuthClient(conn), nil
}
