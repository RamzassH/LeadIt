package facade

import (
	"context"

	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/employee"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/organization"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/project"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/role"
)

type Facade struct {
	org  organization.Organization
	role role.Role
	proj project.Project
	emp  employee.Employee
}

func NewFacade(org organization.Organization, role role.Role, proj project.Project, emp employee.Employee) *Facade {
	return &Facade{
		org:  org,
		role: role,
		proj: proj,
		emp:  emp,
	}
}

func (f *Facade) AddOrganization(ctx context.Context, payload models.AddOrganizationPayload) (int64, error) {
	return f.org.AddOrganization(ctx, payload)
}

func (f *Facade) GetOrganization(ctx context.Context, id int64) (*models.Organization, error) {
	return f.org.GetOrganization(ctx, id)
}

func (f *Facade) GetAllOrganizations(ctx context.Context) ([]models.Organization, error) {
	return f.org.GetAllOrganizations(ctx)
}

func (f *Facade) UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (int64, error) {
	return f.org.UpdateOrganization(ctx, payload)
}

func (f *Facade) DeleteOrganization(ctx context.Context, id int64) (int64, error) {
	return f.org.DeleteOrganization(ctx, id)
}

func (f *Facade) AddRole(ctx context.Context, payload models.AddRolePayload) (int64, error) {
	return f.role.AddRole(ctx, payload)
}

func (f *Facade) GetRole(ctx context.Context, id int64) (*models.Role, error) {
	return f.role.GetRole(ctx, id)
}

func (f *Facade) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	return f.role.GetAllRoles(ctx)
}

func (f *Facade) UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (int64, error) {
	return f.role.UpdateRole(ctx, payload)
}

func (f *Facade) DeleteRole(ctx context.Context, id int64) (int64, error) {
	return f.role.DeleteRole(ctx, id)
}

func (f *Facade) AddProject(ctx context.Context, payload models.AddProjectPayload) (int64, error) {
	return f.proj.AddProject(ctx, payload)
}

func (f *Facade) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	return f.proj.GetProject(ctx, id)
}

func (f *Facade) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	return f.proj.GetAllProjects(ctx)
}

func (f *Facade) UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (int64, error) {
	return f.proj.UpdateProject(ctx, payload)
}

func (f *Facade) DeleteProject(ctx context.Context, id int64) (int64, error) {
	return f.proj.DeleteProject(ctx, id)
}

// Методы для работы с сотрудниками
func (f *Facade) AddEmployee(ctx context.Context, payload models.AddEmployee) (int64, error) {
	return f.emp.AddEmployee(ctx, payload)
}

func (f *Facade) GetEmployee(ctx context.Context, id int64) (*models.Employee, error) {
	return f.emp.GetEmployee(ctx, id)
}

func (f *Facade) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	return f.emp.GetAllEmployees(ctx)
}

func (f *Facade) UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (int64, error) {
	return f.emp.UpdateEmployee(ctx, payload)
}

func (f *Facade) DeleteEmployee(ctx context.Context, id int64) (int64, error) {
	return f.emp.DeleteEmployee(ctx, id)
}
