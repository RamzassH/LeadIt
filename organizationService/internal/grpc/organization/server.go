package organization

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Organization interface {
}

type ServerAPI struct {
	organizationv1.UnimplementedOrganizationServer
	organization Organization
	logger       zerolog.Logger
	validate     *validator.Validate
}

func RegisterGRPCServer(grpcServer *grpc.Server, validate *validator.Validate, logger zerolog.Logger, organizationService Organization) {
	organizationv1.RegisterOrganizationServer(grpcServer, &ServerAPI{
		validate:     validate,
		organization: organizationService,
		logger:       logger,
	})
}
