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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (models.Organization, error)

	DeleteOrganization(ctx context.Context, id int64) (int64, error)
}

type Role interface {
	AddRole(ctx context.Context, payload models.AddRolePayload) (int64, error)
	GetRole(ctx context.Context, id int64) (*models.Role, error)
	GetAllRoles(ctx context.Context) ([]models.Role, error)
	UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (int64, error)
	DeleteRole(ctx context.Context, id int64) (int64, error)
}

type Project interface {
	AddProject(ctx context.Context, payload models.AddProjectPayload) (int64, error)
	GetProject(ctx context.Context, id int64) (*models.Project, error)
	GetAllProjects(ctx context.Context) ([]models.Project, error)
	UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (int64, error)
	DeleteProject(ctx context.Context, id int64) (int64, error)
}

type Employee interface {
	AddEmployee(ctx context.Context, payload models.AddEmployee) (int64, error)
	GetEmployee(ctx context.Context, id int64) (*models.Employee, error)
	GetAllEmployees(ctx context.Context) ([]models.Employee, error)
	UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (int64, error)
	DeleteEmployee(ctx context.Context, id int64) (int64, error)
}

type ServerAPI struct {
	organizationv1.UnimplementedOrganizationServer
	rolev1.UnimplementedRoleServer
	employeev1.UnimplementedEmployeeServer
	projectv1.UnimplementedProjectServer
	organization Organization
	role         Role
	employee     Employee
	project      Project
	logger       zerolog.Logger
	validate     *validator.Validate
}

func RegisterGRPCServer(grpcServer *grpc.Server, validate *validator.Validate, logger zerolog.Logger, organizationService Service) {
	organizationv1.RegisterOrganizationServer(grpcServer, &ServerAPI{
		validate:     validate,
		organization: organizationService,
		logger:       logger,
	})
}

func (s *ServerAPI) ValidateStruct(data interface{}) error {
	if err := s.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func (s *ServerAPI) AddOrganization(ctx context.Context, req *organizationv1.AddOrganizationRequest) (*organizationv1.AddOrganizationResponse, error) {
	addOrganizationReq := models.AddOrganizationPayload{
		Name:              req.GetName(),
		Description:       req.GetDescription(),
		OrganizationImage: req.GetImage(),
	}

	if err := s.ValidateStruct(addOrganizationReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	OrganizationId, err := s.organization.AddOrganization(ctx, addOrganizationReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &organizationv1.AddOrganizationResponse{
		Id: OrganizationId,
	}, nil
}

func (s *ServerAPI) GetOrganization(ctx context.Context, req *organizationv1.GetOrganizationRequest) (*organizationv1.GetOrganizationResponse, error) {
	getOrganizationReq := models.GetOrganizationPayload{
		OrganizationID: req.GetId(),
	}

	if err := s.ValidateStruct(getOrganizationReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	organization, err := s.organization.GetOrganization(ctx, getOrganizationReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if organization == nil {
		return &organizationv1.GetOrganizationResponse{}, nil
	}

	orgResponse := &organizationv1.OrganizationType{
		Id:                organization.ID,
		Name:              organization.Name,
		Description:       organization.Description,
		OrganizerId:       organization.OrganizerID,
		OrganizationImage: organization.OrganizationImage,
	}

	return &organizationv1.GetOrganizationResponse{
		Organization: orgResponse,
	}, nil
}

func (s *ServerAPI) GetOrganizations(ctx context.Context, req *organizationv1.GetOrganizationsRequest) (*organizationv1.GetOrganizationsResponse, error) {
	getOrganizationsReq := models.GetOrganizationsPayload{
		OrganizerID: req.GetOrganizerId(),
	}

	if err := s.ValidateStruct(getOrganizationsReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	organizations, err := s.organization.GetAllOrganizations(ctx, getOrganizationsReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get organizations: %v", err)
	}

	var orgResponse []*organizationv1.OrganizationType

	for _, organization := range organizations {
		orgResponse = append(orgResponse, &organizationv1.OrganizationType{
			Id:                organization.ID,
			Name:              organization.Name,
			Description:       organization.Description,
			OrganizerId:       organization.OrganizerID,
			OrganizationImage: organization.OrganizationImage,
		})
	}

	return &organizationv1.GetOrganizationsResponse{
		Organizations: orgResponse,
	}, nil
}

func (s *ServerAPI) UpdateOrganization(ctx context.Context, req *organizationv1.UpdateOrganizationRequest) (*organizationv1.UpdateOrganizationResponse, error) {
	updateOrganizationReq := models.UpdateOrganizationPayload{
		ID:                req.Organization.Id,
		Name:              req.Organization.Name,
		Description:       req.Organization.Description,
		OrganizerID:       req.Organization.OrganizerId,
		OrganizationImage: req.Organization.OrganizationImage,
	}

	if err := s.ValidateStruct(updateOrganizationReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	updatedOrganization, err := s.organization.UpdateOrganization(ctx, updateOrganizationReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	orgResponse := &organizationv1.OrganizationType{
		Id:                updatedOrganization.ID,
		Name:              updatedOrganization.Name,
		Description:       updatedOrganization.Description,
		OrganizerId:       updatedOrganization.OrganizerID,
		OrganizationImage: updatedOrganization.OrganizationImage,
	}

	return &organizationv1.UpdateOrganizationResponse{
		Organization: orgResponse,
	}, nil
}

func (s *ServerAPI) DeleteOrganization(ctx context.Context, req *organizationv1.DeleteOrganizationRequest) (*organizationv1.DeleteOrganizationResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "organization ID is required")
	}

	deletedID, err := s.organization.DeleteOrganization(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete organization: %v", err)
	}

	return &organizationv1.DeleteOrganizationResponse{Id: deletedID}, nil
}

func (s *ServerAPI) AddRole(ctx context.Context, req *rolev1.AddRoleRequest) (*rolev1.AddRoleResponse, error) {
	payload := models.AddRolePayload{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	roleID, err := s.role.AddRole(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add role: %v", err)
	}

	return &rolev1.AddRoleResponse{Id: roleID}, nil
}
func (s *ServerAPI) GetRole(ctx context.Context, req *rolev1.GetRoleRequest) (*rolev1.GetRoleResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "role ID is required")
	}

	role, err := s.role.GetRole(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "role not found: %v", err)
	}

	return &rolev1.GetRoleResponse{
		Role: &rolev1.RoleType{
			Id:          role.ID,
			Name:        role.Name,
			Permissions: role.Permissions,
		},
	}, nil
}
func (s *ServerAPI) GetRoles(ctx context.Context, req *rolev1.GetRolesRequest) (*rolev1.GetRolesResponse, error) {
	roles, err := s.role.GetAllRoles(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get roles: %v", err)
	}

	var response []*rolev1.RoleType
	for _, r := range roles {
		response = append(response, &rolev1.RoleType{
			Id:          r.ID,
			Name:        r.Name,
			Permissions: r.Permissions,
		})
	}

	return &rolev1.GetRolesResponse{Roles: response}, nil
}
func (s *ServerAPI) UpdateRole(ctx context.Context, req *rolev1.UpdateRoleRequest) (*rolev1.UpdateRoleResponse, error) {
	payload := models.UpdateRolePayload{
		ID:             req.Role.Id,
		Name:           req.Role.Name,
		OrganizationID: req.Role.OrganizationId,
		Permissions:    req.Role.Permissions,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedID, err := s.role.UpdateRole(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update role: %v", err)
	}

	return &rolev1.UpdateRoleResponse{
		Role: &rolev1.RoleType{
			Id:          updatedID,
			Name:        req.Role.Name,
			Permissions: req.Role.Permissions,
		},
	}, nil
}
func (s *ServerAPI) DeleteRole(ctx context.Context, req *rolev1.DeleteRoleRequest) (*rolev1.DeleteRoleResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "role ID is required")
	}

	deletedID, err := s.role.DeleteRole(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete role: %v", err)
	}

	return &rolev1.DeleteRoleResponse{Id: deletedID}, nil
}

func (s *ServerAPI) AddProject(ctx context.Context, req *projectv1.AddProjectRequest) (*projectv1.AddProjectResponse, error) {
	payload := models.AddProjectPayload{
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		OrganizationID: req.GetOrganizationId(),
		Image:          req.GetImage(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	projectID, err := s.project.AddProject(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add project: %v", err)
	}

	return &projectv1.AddProjectResponse{Id: projectID}, nil
}
func (s *ServerAPI) GetProject(ctx context.Context, req *projectv1.GetProjectRequest) (*projectv1.GetProjectResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "project ID is required")
	}

	project, err := s.project.GetProject(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project not found: %v", err)
	}

	return &projectv1.GetProjectResponse{
		Project: &projectv1.ProjectType{
			Id:             project.ID,
			Name:           project.Name,
			Description:    project.Description,
			OrganizationId: project.OrganizationID,
			Image:          project.Image,
		},
	}, nil
}
func (s *ServerAPI) GetProjects(ctx context.Context, req *projectv1.GetProjectsRequest) (*projectv1.GetProjectsResponse, error) {
	projects, err := s.project.GetAllProjects(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get projects: %v", err)
	}

	var response []*projectv1.ProjectType
	for _, project := range projects {
		response = append(response, &projectv1.ProjectType{
			Id:             project.ID,
			Name:           project.Name,
			Description:    project.Description,
			OrganizationId: project.OrganizationID,
			Image:          project.Image,
		})
	}

	return &projectv1.GetProjectsResponse{
		Projects: response}, nil
}

func (s *ServerAPI) UpdateProject(ctx context.Context, req *projectv1.UpdateProjectRequest) (*projectv1.UpdateProjectResponse, error) {
	payload := models.UpdateProjectPayload{
		ID:             req.Project.Id,
		Name:           req.Project.Name,
		Description:    req.Project.Description,
		OrganizationID: req.Project.OrganizationId,
		Image:          req.Project.Image,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedID, err := s.project.UpdateProject(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update project: %v", err)
	}

	return &projectv1.UpdateProjectResponse{
		Project: &projectv1.ProjectType{
			Id:             updatedID,
			Name:           req.Project.Name,
			Description:    req.Project.Description,
			OrganizationId: req.Project.OrganizationId,
			Image:          req.Project.Image,
		},
	}, nil
}
func (s *ServerAPI) DeleteProject(ctx context.Context, req *projectv1.DeleteProjectRequest) (*projectv1.DeleteProjectResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "project ID is required")
	}

	deletedID, err := s.project.DeleteProject(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete project: %v", err)
	}

	return &projectv1.DeleteProjectResponse{Id: deletedID}, nil
}

func (s *ServerAPI) AddEmployee(ctx context.Context, req *employeev1.AddEmployeeRequest) (*employeev1.AddEmployeeResponse, error) {
	payload := models.AddEmployee{
		UserID:         req.GetUserId(),
		OrganizationID: req.GetOrganizationId(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	employeeID, err := s.employee.AddEmployee(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add employee: %v", err)
	}

	return &employeev1.AddEmployeeResponse{Id: employeeID}, nil
}
func (s *ServerAPI) GetEmployee(ctx context.Context, req *employeev1.GetEmployeeRequest) (*employeev1.GetEmployeeResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "employee ID is required")
	}

	employee, err := s.employee.GetEmployee(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "employee not found: %v", err)
	}

	return &employeev1.GetEmployeeResponse{
		Employee: &employeev1.EmployeeType{
			Id:             employee.ID,
			UserId:         employee.UserID,
			OrganizationId: employee.OrganizationID,
		},
	}, nil
}
func (s *ServerAPI) GetEmployees(ctx context.Context, req *employeev1.GetEmployeesRequest) (*employeev1.GetEmployeesResponse, error) {
	if req.GetOrganizationId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "organization ID is required")
	}

	employees, err := s.employee.GetAllEmployees(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get employees: %v", err)
	}

	var response []*employeev1.EmployeeType
	for _, e := range employees {
		response = append(response, &employeev1.EmployeeType{
			Id:             e.ID,
			UserId:         e.UserID,
			OrganizationId: e.OrganizationID,
		})
	}

	return &employeev1.GetEmployeesResponse{Employees: response}, nil
}
func (s *ServerAPI) UpdateEmployee(ctx context.Context, req *employeev1.UpdateEmployeeRequest) (*employeev1.UpdateEmployeeResponse, error) {
	payload := models.UpdateEmployee{
		ID:             req.Employee.Id,
		UserID:         req.Employee.UserId,
		OrganizationID: req.Employee.OrganizationId,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	_, err := s.employee.UpdateEmployee(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update employee: %v", err)
	}

	return &employeev1.UpdateEmployeeResponse{
		Employee: &employeev1.EmployeeType{
			Id:             payload.ID,
			UserId:         payload.UserID,
			OrganizationId: payload.OrganizationID,
		},
	}, nil
}
func (s *ServerAPI) DeleteEmployee(ctx context.Context, req *employeev1.DeleteEmployeeRequest) (*employeev1.DeleteEmployeeResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "employee ID is required")
	}

	deletedID, err := s.employee.DeleteEmployee(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete employee: %v", err)
	}

	return &employeev1.DeleteEmployeeResponse{Id: deletedID}, nil
}
