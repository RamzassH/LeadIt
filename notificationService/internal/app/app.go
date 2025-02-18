package app

import (
	grpcapp "github.com/RamzassH/LeadIt/notificationService/internal/app/grpc"
	notification "github.com/RamzassH/LeadIt/notificationService/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

func New(
	logger *zerolog.Logger,
	grpcPort int,
	validate *validator.Validate,
	smtpHost, smtpPort, smtpUser, smtpPassword string) (*App, error) {

	SMTP := notification.SMTP{
		Host:     smtpHost,
		Port:     smtpPort,
		User:     smtpUser,
		Password: smtpPassword,
	}
	notificationService := notification.New(
		logger,
		SMTP,
	)

	grpcApp := grpcapp.New(
		logger,
		grpcPort,
		validate,
		notificationService,
	)

	return &App{
		GRPCServer: grpcApp,
		validator:  validate,
	}, nil
}
