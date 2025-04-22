package organization

import (
	"context"
	organizationv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/organization"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/grpc/interceptors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerAPI) CreateOrganization(ctx context.Context, req *organizationv1.CreateOrganizationRequest) (*organizationv1.CreateOrganizationResponse, error) {
	addOrganizationReq := models.CreateOrganizationDTO{
		Name:              req.GetName(),
		Description:       req.GetDescription(),
		OrganizationImage: req.GetImage(),
	}

	if err := s.ValidateStruct(addOrganizationReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	OrganizationId, err := s.service.CreateOrganization(ctx, addOrganizationReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &organizationv1.CreateOrganizationResponse{
		Id: OrganizationId,
	}, nil
}

func (s *ServerAPI) GetOrganization(ctx context.Context, req *organizationv1.GetOrganizationRequest) (*organizationv1.GetOrganizationResponse, error) {
	getOrganizationReq := models.GetOrganizationDTO{
		OrganizationID: req.GetId(),
	}

	if err := s.ValidateStruct(getOrganizationReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	organization, err := s.service.GetOrganization(ctx, getOrganizationReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if organization == nil {
		return &organizationv1.GetOrganizationResponse{}, nil
	}

	orgResponse := &organizationv1.OrganizationType{
		Id:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		OrganizerId: organization.OrganizerID,
		Image:       organization.Image,
	}

	return &organizationv1.GetOrganizationResponse{
		Organization: orgResponse,
	}, nil
}

func (s *ServerAPI) GetOrganizations(ctx context.Context, req *organizationv1.GetOrganizationsRequest) (*organizationv1.GetOrganizationsResponse, error) {

	userIDValue := ctx.Value(interceptors.CtxUserID)
	if userIDValue == nil {
		return nil, status.Errorf(codes.Unauthenticated, "userID not found in context")
	}

	userID := userIDValue.(int64)
	getOrganizationsReq := models.GetOrganizationsDTO{
		OrganizerID: userID,
	}

	if err := s.ValidateStruct(getOrganizationsReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	organizations, err := s.service.GetAllOrganizations(ctx, getOrganizationsReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get organizations: %v", err)
	}

	var orgResponse []*organizationv1.OrganizationType

	for _, organization := range organizations {
		orgResponse = append(orgResponse, &organizationv1.OrganizationType{
			Id:          organization.ID,
			Name:        organization.Name,
			Description: organization.Description,
			OrganizerId: organization.OrganizerID,
			Image:       organization.Image,
		})
	}

	return &organizationv1.GetOrganizationsResponse{
		Organizations: orgResponse,
	}, nil
}

func (s *ServerAPI) UpdateOrganization(ctx context.Context, req *organizationv1.UpdateOrganizationRequest) (*organizationv1.UpdateOrganizationResponse, error) {
	updateOrganizationReq := models.UpdateOrganizationDTO{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
	}

	if err := s.ValidateStruct(updateOrganizationReq); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	updatedOrganization, err := s.service.UpdateOrganization(ctx, updateOrganizationReq)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	orgResponse := &organizationv1.OrganizationType{
		Id:          updatedOrganization.ID,
		Name:        updatedOrganization.Name,
		Description: updatedOrganization.Description,
		OrganizerId: updatedOrganization.OrganizerID,
		Image:       updatedOrganization.Image,
	}

	return &organizationv1.UpdateOrganizationResponse{
		Organization: orgResponse,
	}, nil
}

func (s *ServerAPI) DeleteOrganization(ctx context.Context, req *organizationv1.DeleteOrganizationRequest) (*organizationv1.DeleteOrganizationResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "organization ID is required")
	}

	deletedID, err := s.service.DeleteOrganization(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete organization: %v", err)
	}

	return &organizationv1.DeleteOrganizationResponse{Id: deletedID}, nil
}
