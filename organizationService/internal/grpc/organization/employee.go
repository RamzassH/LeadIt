package organization

import (
	"context"
	employeev1 "github.com/RamzassH/LeadIt/libs/contracts/gen/employee"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) AddEmployee(ctx context.Context, req *employeev1.AddEmployeeRequest) (*employeev1.AddEmployeeResponse, error) {
	payload := models.AddEmployee{
		UserID:         req.GetUserId(),
		OrganizationID: req.GetOrganizationId(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	employeeID, err := s.service.AddEmployee(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add employee: %v", err)
	}

	return &employeev1.AddEmployeeResponse{Id: employeeID}, nil
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
func (s *ServerAPI) UpdateEmployee(ctx context.Context, req *employeev1.UpdateEmployeeRequest) (*employeev1.UpdateEmployeeResponse, error) {
	payload := models.UpdateEmployee{
		ID:             req.Employee.Id,
		UserID:         req.Employee.UserId,
		OrganizationID: req.Employee.OrganizationId,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	_, err := s.service.UpdateEmployee(ctx, payload)
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

	deletedID, err := s.service.DeleteEmployee(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete employee: %v", err)
	}

	return &employeev1.DeleteEmployeeResponse{Id: deletedID}, nil
}
