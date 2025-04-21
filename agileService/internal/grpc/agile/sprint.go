package agile

import (
	"context"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	sprintV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/sprint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServerApi) CreateSprint(ctx context.Context, req *sprintV1.CreateSprintRequest) (*sprintV1.CreateSprintResponse, error) {
	payload := models.CreateSprintDTO{
		ProjectID:   req.GetProjectId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		StartDate:   req.GetStartDate().AsTime(),
		EndDate:     req.GetEndDate().AsTime(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id, err := s.service.CreateSprint(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create sprint: %v", err)
	}

	return &sprintV1.CreateSprintResponse{Id: id}, nil
}

func (s *ServerApi) GetSprint(ctx context.Context, req *sprintV1.GetSprintRequest) (*sprintV1.SprintType, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid sprint id")
	}

	sprint, err := s.service.GetSprint(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get sprint: %v", err)
	}

	return &sprintV1.SprintType{
		Id:          sprint.ID,
		ProjectId:   sprint.ProjectID,
		Name:        sprint.Name,
		Description: sprint.Description,
		StartDate:   timestamppb.New(sprint.StartDate),
		EndDate:     timestamppb.New(sprint.EndDate),
		IsActive:    sprint.IsActive,
	}, nil
}

func (s *ServerApi) GetSprintsByProject(ctx context.Context, req *sprintV1.GetSprintsByProjectRequest) (*sprintV1.GetSprintsByProjectResponse, error) {
	projectId := req.GetProjectId()
	if projectId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid project id")
	}

	sprints, err := s.service.GetSprintsByProject(ctx, projectId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get sprints: %v", err)
	}

	resp := make([]*sprintV1.SprintType, 0, len(sprints))
	for _, sprint := range sprints {
		resp = append(resp, &sprintV1.SprintType{
			Id:          sprint.ID,
			ProjectId:   sprint.ProjectID,
			Name:        sprint.Name,
			Description: sprint.Description,
			StartDate:   timestamppb.New(sprint.StartDate),
			EndDate:     timestamppb.New(sprint.EndDate),
			IsActive:    sprint.IsActive,
		})
	}

	return &sprintV1.GetSprintsByProjectResponse{Sprints: resp}, nil
}

func (s *ServerApi) UpdateSprint(ctx context.Context, req *sprintV1.UpdateSprintRequest) (*sprintV1.SprintType, error) {
	payload := models.UpdateSprintDTO{
		ID:          req.GetId(),
		Name:        req.Name,
		Description: req.Description,
		StartDate:   req.GetStartDate().AsTime(),
		EndDate:     req.GetEndDate().AsTime(),
		IsActive:    req.IsActive,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	sprint, err := s.service.UpdateSprint(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update sprint: %v", err)
	}

	return &sprintV1.SprintType{
		Id:          sprint.ID,
		ProjectId:   sprint.ProjectID,
		Name:        sprint.Name,
		Description: sprint.Description,
		StartDate:   timestamppb.New(sprint.StartDate),
		EndDate:     timestamppb.New(sprint.EndDate),
		IsActive:    sprint.IsActive,
	}, nil
}

func (s *ServerApi) DeleteSprint(ctx context.Context, req *sprintV1.DeleteSprintRequest) (*sprintV1.DeleteSprintResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid sprint id")
	}

	deletedId, err := s.service.DeleteSprint(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete sprint: %v", err)
	}

	return &sprintV1.DeleteSprintResponse{Id: deletedId}, nil
}

func (s *ServerApi) SetSprintIsActive(ctx context.Context, req *sprintV1.SetSprintIsActiveRequest) (*sprintV1.SetSprintIsActiveResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid sprint id")
	}

	err := s.service.SetSprintIsActive(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update is_active: %v", err)
	}

	return &sprintV1.SetSprintIsActiveResponse{Success: true}, nil
}
