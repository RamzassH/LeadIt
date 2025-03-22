package organization

import (
	"context"
	projectv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/project"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	projectID, err := s.service.AddProject(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add project: %v", err)
	}

	return &projectv1.AddProjectResponse{Id: projectID}, nil
}
func (s *ServerAPI) GetProject(ctx context.Context, req *projectv1.GetProjectRequest) (*projectv1.GetProjectResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "project ID is required")
	}

	project, err := s.service.GetProject(ctx, req.GetId())
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
	organizationId := req.GetOrganizationId()
	if organizationId == 0 {
		return nil, status.Error(codes.InvalidArgument, "organization ID is required")
	}

	projects, err := s.service.GetAllProjects(ctx, organizationId)
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

	updatedProject, err := s.service.UpdateProject(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update project: %v", err)
	}

	return &projectv1.UpdateProjectResponse{
		Project: &projectv1.ProjectType{
			Id:             updatedProject.ID,
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

	deletedID, err := s.service.DeleteProject(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete project: %v", err)
	}

	return &projectv1.DeleteProjectResponse{Id: deletedID}, nil
}
