package task

import (
	"context"
	actionv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/action"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s ServerAPI) CreateAction(ctx context.Context, req *actionv1.CreateActionRequest) (*actionv1.CreateActionResponse, error) {
	payload := models.CreateActionDTO{
		TaskID:      req.GetTaskId(),
		UserID:      req.GetUserId(),
		ActionType:  req.GetActionType(),
		Description: req.GetDescription(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	actionID, err := s.service.CreateAction(ctx, payload)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add action: %v", err)
	}

	return &actionv1.CreateActionResponse{
		Id: actionID,
	}, nil
}

func (s ServerAPI) GetActionsByTask(ctx context.Context, req *actionv1.GetActionsByTaskRequest) (*actionv1.GetActionsByTaskResponse, error) {
	taskID := req.GetTaskId()
	if taskID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	actions, err := s.service.GetActionsForTask(ctx, taskID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get actions: %v", err)
	}

	var response []*actionv1.ActionType

	for _, action := range actions {
		response = append(response, &actionv1.ActionType{
			Id:          action.ID,
			UserId:      action.UserID,
			TaskId:      action.TaskID,
			Description: action.Description,
		})
	}

	return &actionv1.GetActionsByTaskResponse{
		Actions: response,
	}, nil
}
