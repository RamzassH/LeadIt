package task

import (
	"context"
	statusv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/status"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s ServerAPI) CreateStatus(ctx context.Context, req *statusv1.CreateStatusRequest) (*statusv1.CreateStatusResponse, error) {
	payload := models.CreateStatusDTO{
		ProgressBarID:  req.GetProgressBarId(),
		Name:           req.GetName(),
		Order:          req.GetOrder(),
		NotifyRolesIDS: req.GetNotifyRolesIds(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id, err := s.service.CreateStatus(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create status: %s", err.Error())
	}

	return &statusv1.CreateStatusResponse{Id: id}, nil
}

func (s ServerAPI) GetManyStatusesByProgressBar(ctx context.Context, req *statusv1.GetStatusesByProgressBarRequest) (*statusv1.GetStatusesByProgressBarResponse, error) {
	pbID := req.GetProgressBarId()
	if pbID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid progress bar id")
	}

	statuses, err := s.service.GetStatusForProgressBar(ctx, pbID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get statuses: %s", err.Error())
	}

	return &statusv1.GetStatusesByProgressBarResponse{
		Statuses: []*statusv1.StatusType{
			{
				Id:             statuses.ID,
				ProgressBarId:  statuses.ProgressBarID,
				Name:           statuses.Name,
				Order:          statuses.Order,
				NotifyRolesIds: statuses.NotifyRolesIDS,
			},
		},
	}, nil
}

func (s ServerAPI) UpdateStatus(ctx context.Context, req *statusv1.UpdateStatusRequest) (*statusv1.UpdateStatusResponse, error) {
	payload := models.UpdateStatusDTO{
		ID:             req.GetId(),
		Name:           req.GetName(),
		Order:          req.GetOrder(),
		NotifyRolesIDS: req.GetNotifyRolesIds(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedStatus, err := s.service.UpdateStatus(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update status: %s", err.Error())
	}

	return &statusv1.UpdateStatusResponse{
		UpdatedStatus: &statusv1.StatusType{
			Id:             updatedStatus.ID,
			ProgressBarId:  updatedStatus.ProgressBarID,
			Name:           updatedStatus.Name,
			Order:          updatedStatus.Order,
			NotifyRolesIds: updatedStatus.NotifyRolesIDS,
		},
	}, nil
}

func (s ServerAPI) DeleteStatus(ctx context.Context, req *statusv1.DeleteStatusRequest) (*statusv1.DeleteStatusResponse, error) {
	id := req.GetId()
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid status id")
	}

	deletedID, err := s.service.DeleteStatus(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete status: %s", err.Error())
	}

	return &statusv1.DeleteStatusResponse{Id: deletedID}, nil
}
