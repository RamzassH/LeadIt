package organization

import (
	"context"
	employeev1 "github.com/RamzassH/LeadIt/libs/contracts/gen/employee"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) CreateEmployee(ctx context.Context, req *employeev1.CreateEmployeeRequest) (*employeev1.CreateEmployeeResponse, error) {
	payload := models.CreateEmployeeDTO{
		UserID:         req.GetUserId(),
		OrganizationID: req.GetOrganizationId(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	employeeID, err := s.service.CreateEmployee(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add employee: %v", err)
	}

	return &employeev1.CreateEmployeeResponse{Id: employeeID}, nil
}
func (s *ServerAPI) GetEmployee(ctx context.Context, req *employeev1.GetEmployeeRequest) (*employeev1.GetEmployeeResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "employee ID is required")
	}

	employee, err := s.service.GetEmployee(ctx, req.GetId())
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
	organizationId := req.GetOrganizationId()
	if organizationId == 0 {
		return nil, status.Error(codes.InvalidArgument, "organization ID is required")
	}

	employees, err := s.service.GetAllEmployees(ctx, organizationId)
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

func (s *ServerAPI) UpdateEmployeeRole(ctx context.Context, req *employeev1.UpdateEmployeeRoleRequest) (*employeev1.UpdateEmployeeRoleResponse, error) {
	payload := models.UpdateEmployeeRoleDTO{
		ID:     req.GetId(),
		RoleID: req.RoleId,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	_, err := s.service.UpdateEmployeeRole(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update employee: %v", err)
	}

	return &employeev1.UpdateEmployeeRoleResponse{
		Id: payload.ID,
	}, nil
}
func (s *ServerAPI) DeleteEmployee(ctx context.Context, req *employeev1.DeleteEmployeeRequest) (*employeev1.DeleteEmployeeResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "employee ID is required")
	}

	deletedID, err := s.service.DeleteEmployee(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete employee: %v", err)
	}

	return &employeev1.DeleteEmployeeResponse{Id: deletedID}, nil
}
