package task

import (
	"context"
	taskv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/task"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s ServerAPI) CrateTask(ctx context.Context, req *taskv1.CreateTaskRequest) (*taskv1.CreateTaskResponse, error) {
	payload := models.CreateTaskDTO{
		Name:            req.GetName(),
		Description:     req.GetDescription(),
		ProjectID:       req.GetProjectId(),
		SetterID:        req.GetSetterId(),
		SolverID:        req.GetSolverId(),
		ProgressBarID:   req.GetProgressBarId(),
		CurrentStatusID: req.GetCurrentStatusId(),
		AllocatedTime:   req.GetAllocatedTime(),
		SprintID:        req.GetSprintId(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	taskID, err := s.service.CreateTask(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create task: %s", err.Error())
	}

	return &taskv1.CreateTaskResponse{Id: taskID}, nil
}

func (s ServerAPI) GetTask(ctx context.Context, req *taskv1.GetTaskRequest) (*taskv1.GetTaskResponse, error) {
	taskID := req.GetId()
	if taskID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	task, err := s.service.GetTask(ctx, taskID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get task: %s", err.Error())
	}

	return &taskv1.GetTaskResponse{
		Task: &taskv1.TaskType{
			Id:              task.ID,
			ProjectId:       task.ProjectID,
			SetterId:        task.SetterID,
			SolverId:        task.SolverID,
			ProgressBarId:   task.ProgressBarID,
			CurrentStatusId: task.CurrentStatusID,
			AllocatedTime:   task.AllocatedTime,
			WastedTime:      task.WastedTime,
			LastStartTime:   timestamppb.New(task.LastStartTime),
		},
	}, nil
}

func (s ServerAPI) GetTasksByProject(ctx context.Context, req *taskv1.GetTasksByProjectRequest) (*taskv1.GetTasksByProjectResponse, error) {
	projectID := req.GetProjectId()
	if projectID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid project id")
	}

	tasks, err := s.service.GetTasksForProject(ctx, projectID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get tasks: %s", err.Error())
	}

	resp := make([]*taskv1.TaskType, 0, len(tasks))
	for _, task := range tasks {
		resp = append(resp, &taskv1.TaskType{
			Id:              task.ID,
			ProjectId:       task.ProjectID,
			SetterId:        task.SetterID,
			SolverId:        task.SolverID,
			ProgressBarId:   task.ProgressBarID,
			CurrentStatusId: task.CurrentStatusID,
			AllocatedTime:   task.AllocatedTime,
			WastedTime:      task.WastedTime,
			LastStartTime:   timestamppb.New(task.LastStartTime),
		})
	}

	return &taskv1.GetTasksByProjectResponse{Tasks: resp}, nil
}

func (s ServerAPI) UpdateTask(ctx context.Context, req *taskv1.UpdateTaskRequest) (*taskv1.UpdateTaskResponse, error) {
	payload := models.UpdateTaskDTO{
		ID:              req.GetId(),
		Name:            req.GetName(),
		Description:     req.GetDescription(),
		SprintID:        req.GetSprintId(),
		SolverID:        req.GetSolverId(),
		ProgressBarID:   req.GetProgressBarId(),
		CurrentStatusID: req.GetCurrentStatusId(),
		AllocatedTime:   req.GetAllocatedTime(),
		WastedTime:      req.GetWastedTime(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedTask, err := s.service.UpdateTask(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update task: %s", err.Error())
	}

	return &taskv1.UpdateTaskResponse{
		Task: &taskv1.TaskType{
			Id:              updatedTask.ID,
			ProjectId:       updatedTask.ProjectID,
			SetterId:        updatedTask.SetterID,
			SolverId:        updatedTask.SolverID,
			ProgressBarId:   updatedTask.ProgressBarID,
			CurrentStatusId: updatedTask.CurrentStatusID,
			AllocatedTime:   updatedTask.AllocatedTime,
			WastedTime:      updatedTask.WastedTime,
			LastStartTime:   timestamppb.New(updatedTask.LastStartTime),
		},
	}, nil
}

func (s ServerAPI) DeleteTask(ctx context.Context, req *taskv1.DeleteTaskRequest) (*taskv1.DeleteTaskResponse, error) {
	taskID := req.GetId()
	if taskID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	deletedID, err := s.service.DeleteTask(ctx, taskID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete task: %s", err.Error())
	}

	return &taskv1.DeleteTaskResponse{Id: deletedID}, nil
}
