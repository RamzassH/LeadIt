package organization

import (
	"context"
	rolev1 "github.com/RamzassH/LeadIt/libs/contracts/gen/role"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) AddRole(ctx context.Context, req *rolev1.AddRoleRequest) (*rolev1.AddRoleResponse, error) {
	payload := models.AddRolePayload{
		Name:           req.GetName(),
		OrganizationID: req.GetOrganizationId(),
		Permissions:    req.GetPermissions(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	roleID, err := s.service.AddRole(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add role: %v", err)
	}

	return &rolev1.AddRoleResponse{Id: roleID}, nil
}
func (s *ServerAPI) GetRole(ctx context.Context, req *rolev1.GetRoleRequest) (*rolev1.GetRoleResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "role ID is required")
	}

	role, err := s.service.GetRole(ctx, req.GetId())
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
	organizationId := req.GetOrganizationId()
	if organizationId == 0 {
		return nil, status.Error(codes.InvalidArgument, "organization ID is required")
	}

	roles, err := s.service.GetAllRoles(ctx, organizationId)
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

	updatedRole, err := s.service.UpdateRole(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update role: %v", err)
	}

	return &rolev1.UpdateRoleResponse{
		Role: &rolev1.RoleType{
			Id:             updatedRole.ID,
			OrganizationId: updatedRole.OrganizationID,
			Name:           updatedRole.Name,
			Permissions:    updatedRole.Permissions,
		},
	}, nil
}
func (s *ServerAPI) DeleteRole(ctx context.Context, req *rolev1.DeleteRoleRequest) (*rolev1.DeleteRoleResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "role ID is required")
	}

	deletedID, err := s.service.DeleteRole(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete role: %v", err)
	}

	return &rolev1.DeleteRoleResponse{Id: deletedID}, nil
}
