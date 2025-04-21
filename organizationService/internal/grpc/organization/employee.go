package organization

import (
	"context"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	employeev1 "github.com/RamzassH/LeadIt/libs/contracts/gen/employee"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/grpc/interceptors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
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
	organizationIDValue := ctx.Value(interceptors.CtxOrganizationID)
	if organizationIDValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "organizationID not found in context")
	}

	orgStr, ok := organizationIDValue.(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "organizationID type assertion failed")
	}
	organizationID, err := strconv.ParseInt(orgStr, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid organizationID format")
	}

	payload := models.UpdateEmployeeRoleDTO{
		ID:     req.GetId(),
		RoleID: req.RoleId,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedEmployeeId, err := s.service.UpdateEmployeeRole(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update employee: %v", err)
	}

	_, err = s.authClient.ReissueAccessTokenWithContext(ctx, &authv1.ReissueTokenWithContextRequest{
		UserId:         payload.ID,
		OrganizationId: organizationID,
		RoleId:         payload.RoleID,
	})
	if err != nil {
		s.logger.Warn().Err(err).Msg("failed to reissue token after role update")
	}

	return &employeev1.UpdateEmployeeRoleResponse{
		Id: updatedEmployeeId,
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
