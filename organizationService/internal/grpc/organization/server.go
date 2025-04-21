package organization

import (
	"context"
	"fmt"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
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
	CreateOrganization(
		ctx context.Context,
		payload models.CreateOrganizationDTO) (int64, error)

	GetOrganization(ctx context.Context, payload models.GetOrganizationDTO) (*models.OrganizationDTO, error)

	GetAllOrganizations(ctx context.Context, payload models.GetOrganizationsDTO) ([]models.OrganizationDTO, error)

	UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationDTO) (*models.OrganizationDTO, error)
	DeleteOrganization(ctx context.Context, organizationID, organizerID int64) (int64, error)
}

type Role interface {
	CreateRole(ctx context.Context, payload models.CreateRoleDTO) (int64, error)
	GetRole(ctx context.Context, id int64) (*models.RoleDTO, error)
	GetAllRoles(ctx context.Context, organizationId int64) ([]models.RoleDTO, error)
	UpdateRole(ctx context.Context, payload models.UpdateRoleDTO) (*models.RoleDTO, error)
	DeleteRole(ctx context.Context, id int64) (int64, error)
}

type Project interface {
	CreateProject(ctx context.Context, payload models.CreateProjectDTO) (int64, error)
	GetProject(ctx context.Context, id int64) (*models.ProjectDTO, error)
	GetAllProjects(ctx context.Context, organizationId int64) ([]models.ProjectDTO, error)
	UpdateProject(ctx context.Context, payload models.UpdateProjectDTO) (*models.ProjectDTO, error)
	DeleteProject(ctx context.Context, id int64) (int64, error)
}

type Employee interface {
	CreateEmployee(ctx context.Context, payload models.CreateEmployeeDTO) (int64, error)
	GetEmployee(ctx context.Context, id int64) (*models.EmployeeDTO, error)
	GetAllEmployees(ctx context.Context, organizationId int64) ([]models.EmployeeDTO, error)
	UpdateEmployeeRole(ctx context.Context, payload models.UpdateEmployeeRoleDTO) (int64, error)
	DeleteEmployee(ctx context.Context, id int64) (int64, error)
}

type ServerAPI struct {
	organizationv1.UnimplementedOrganizationServer
	rolev1.UnimplementedRoleServer
	employeev1.UnimplementedEmployeeServer
	projectv1.UnimplementedProjectServer

	authClient authv1.AuthClient
	service    Service
	logger     zerolog.Logger
	validate   *validator.Validate
}

func RegisterGRPCServer(
	grpcServer *grpc.Server,
	validate *validator.Validate,
	logger zerolog.Logger,
	service Service,
	authClient authv1.AuthClient,
) {
	server := &ServerAPI{
		validate:   validate,
		logger:     logger,
		service:    service,
		authClient: authClient,
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
