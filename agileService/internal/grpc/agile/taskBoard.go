package agile

import (
	"context"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	taskBoardV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/task_board"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) AttachTaskToBoard(ctx context.Context, req *taskBoardV1.TaskBoardRequest) (*taskBoardV1.TaskBoardResponse, error) {
	payload := models.AttachTaskToBoardDTO{
		TaskID: req.GetTaskId(), BoardID: req.GetBoardId(), SprintID: req.GetSprintId(), ColumnID: req.GetColumnId(), Order: req.GetOrder(),
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	success, err := s.service.AttachTaskToBoard(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "attach task failed: %v", err)
	}
	return &taskBoardV1.TaskBoardResponse{Success: success}, nil
}

func (s *ServerApi) MoveTask(ctx context.Context, req *taskBoardV1.MoveTaskRequest) (*taskBoardV1.MoveTaskResponse, error) {
	payload := models.MoveTaskDTO{
		TaskID: req.GetTaskId(), FromColumnID: req.GetFromColumnId(), ToColumnID: req.GetToColumnId(), NewOrder: req.GetNewOrder(),
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	success, err := s.service.MoveTask(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "move task failed: %v", err)
	}
	return &taskBoardV1.MoveTaskResponse{Success: success}, nil
}

func (s *ServerApi) DetachTaskFromBoard(ctx context.Context, req *taskBoardV1.TaskBoardRequest) (*taskBoardV1.TaskBoardResponse, error) {
	payload := models.TaskBoardDTO{
		TaskID: req.GetTaskId(), BoardID: req.GetBoardId(), SprintID: req.GetSprintId(), ColumnID: req.GetColumnId(), Order: req.GetOrder(),
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	success, err := s.service.DetachTaskFromBoard(ctx, payload)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "detach task failed: %v", err)
	}

	return &taskBoardV1.TaskBoardResponse{Success: success}, nil
}
