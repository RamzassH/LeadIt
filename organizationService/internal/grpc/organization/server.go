package organization

import (
	"context"
	"fmt"
	employeev1 "github.com/RamzassH/LeadIt/libs/contracts/gen/employee"
	organizationv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/organization"
	projectv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/project"
	rolev1 "github.com/RamzassH/LeadIt/libs/contracts/gen/role"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Service interface {
	Organization
	Project
	Role
	Employee
}

type Organization interface {
	AddOrganization(
		ctx context.Context,
		payload models.AddOrganizationPayload) (int64, error)

	GetOrganization(ctx context.Context, payload models.GetOrganizationPayload) (*models.Organization, error)

	GetAllOrganizations(ctx context.Context, payload models.GetOrganizationsPayload) ([]models.Organization, error)

	UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (*models.Organization, error)

	DeleteOrganization(ctx context.Context, id int64) (int64, error)
}

type Role interface {
	AddRole(ctx context.Context, payload models.AddRolePayload) (int64, error)
	GetRole(ctx context.Context, id int64) (*models.Role, error)
	GetAllRoles(ctx context.Context, organizationId int64) ([]models.Role, error)
	UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (*models.Role, error)
	DeleteRole(ctx context.Context, id int64) (int64, error)
}

type Project interface {
	AddProject(ctx context.Context, payload models.AddProjectPayload) (int64, error)
	GetProject(ctx context.Context, id int64) (*models.Project, error)
	GetAllProjects(ctx context.Context, organizationId int64) ([]models.Project, error)
	UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (*models.Project, error)
	DeleteProject(ctx context.Context, id int64) (int64, error)
}

type Employee interface {
	AddEmployee(ctx context.Context, payload models.AddEmployee) (int64, error)
	GetEmployee(ctx context.Context, id int64) (*models.Employee, error)
	GetAllEmployees(ctx context.Context, organizationId int64) ([]models.Employee, error)
	UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, id int64) (int64, error)
}

type ServerAPI struct {
	organizationv1.UnimplementedOrganizationServer
	rolev1.UnimplementedRoleServer
	employeev1.UnimplementedEmployeeServer
	projectv1.UnimplementedProjectServer

	service  Service
	logger   zerolog.Logger
	validate *validator.Validate
}

func RegisterGRPCServer(
	grpcServer *grpc.Server,
	validate *validator.Validate,
	logger zerolog.Logger,
	service Service,
) {
	server := &ServerAPI{
		validate: validate,
		logger:   logger,
		service:  service,
	}

	organizationv1.RegisterOrganizationServer(grpcServer, server)
	rolev1.RegisterRoleServer(grpcServer, server)
	projectv1.RegisterProjectServer(grpcServer, server)
	employeev1.RegisterEmployeeServer(grpcServer, server)
}

func (s *ServerAPI) ValidateStruct(data interface{}) error {
	if err := s.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}
