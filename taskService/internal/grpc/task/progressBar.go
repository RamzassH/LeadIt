package task

import (
	"context"
	progressbarv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/progress_bar"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s ServerAPI) CreateProgressBar(ctx context.Context, req *progressbarv1.CreateProgressBarRequest) (*progressbarv1.CreateProgressBarResponse, error) {
	payload := models.CreateProgressBarDTO{
		Name:      req.GetName(),
		ProjectID: req.GetProjectId(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id, err := s.service.CreateProgressBar(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create progress bar: %s", err.Error())
	}

	return &progressbarv1.CreateProgressBarResponse{Id: id}, nil
}

func (s ServerAPI) GetManyProgressBars(ctx context.Context, req *progressbarv1.GetManyProgressBarsRequest) (*progressbarv1.GetManyProgressBarsResponse, error) {
	projectID := req.GetProjectId()
	if projectID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid project id")
	}

	progressBar, err := s.service.GetProgressBarForProject(ctx, projectID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get progress bars: %s", err.Error())
	}

	return &progressbarv1.GetManyProgressBarsResponse{
		ProgressBarList: []*progressbarv1.ProgressBarType{
			{
				Id:        progressBar.ID,
				Name:      progressBar.Name,
				ProjectId: progressBar.ProjectID,
			},
		},
	}, nil
}

func (s ServerAPI) DeleteProgressBar(ctx context.Context, req *progressbarv1.DeleteProgressBarRequest) (*progressbarv1.DeleteProgressBarResponse, error) {
	id := req.GetId()
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid progress bar id")
	}

	deletedID, err := s.service.DeleteProgressBar(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete progress bar: %s", err.Error())
	}

	return &progressbarv1.DeleteProgressBarResponse{Id: deletedID}, nil
}
