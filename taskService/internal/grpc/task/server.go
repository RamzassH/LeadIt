package task

import (
	"context"
	"fmt"
	actionv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/action"
	commentv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/comment"
	progressbarv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/progress_bar"
	statusv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/status"
	tagv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/tag"
	taskv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/task"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Service interface {
	Action
	Comment
	ProgressBar
	Status
	Tag
	Task
}

type Action interface {
	CreateAction(ctx context.Context, payload models.CreateActionDTO) (int64, error)
	GetActionsForTask(ctx context.Context, taskId int64) ([]*models.ActionDTO, error)
}

type Comment interface {
	CreateComment(ctx context.Context, payload models.CreateCommentDTO) (int64, error)
	GetCommentsForTask(ctx context.Context, taskId int64) ([]*models.CommentDTO, error)
	UpdateComment(ctx context.Context, payload models.UpdateCommentDTO) (*models.CommentDTO, error)
	DeleteComment(ctx context.Context, commentId int64) (int64, error)
}

type ProgressBar interface {
	CreateProgressBar(ctx context.Context, payload models.CreateProgressBarDTO) (int64, error)
	GetProgressBarForProject(ctx context.Context, projectId int64) ([]*models.ProgressBarDTO, error)
	UpdateProgressBar(ctx context.Context, payload models.UpdateProgressBarDTO) (*models.ProgressBarDTO, error)
	DeleteProgressBar(ctx context.Context, progressBarId int64) (int64, error)
}

type Status interface {
	CreateStatus(ctx context.Context, payload models.CreateStatusDTO) (int64, error)
	GetStatusForProgressBar(ctx context.Context, progressBarId int64) (*models.StatusDTO, error)
	GetStatusesForProject(ctx context.Context, projectId int64) ([]*models.StatusDTO, error)
	UpdateStatus(ctx context.Context, payload models.UpdateStatusDTO) (*models.StatusDTO, error)
	DeleteStatus(ctx context.Context, statusId int64) (int64, error)
}
type Tag interface {
	CreateTag(ctx context.Context, payload models.CreateTagDTO) (int64, error)
	GetTagsForProject(ctx context.Context, projectId int64) ([]*models.TagDTO, error)
	UpdateTag(ctx context.Context, payload models.UpdateTagDTO) (*models.TagDTO, error)
	DeleteTag(ctx context.Context, tagId int64) (int64, error)
}
type Task interface {
	CreateTask(ctx context.Context, payload models.CreateTaskDTO) (int64, error)
	GetTask(ctx context.Context, taskId int64) (*models.TaskDTO, error)
	GetTasksForProject(ctx context.Context, projectId int64) ([]*models.TaskDTO, error)
	UpdateTask(ctx context.Context, payload models.UpdateTaskDTO) (*models.TaskDTO, error)
	ChangeStatus(ctx context.Context, payload models.ChangeStatusDTO) (int64, error)
	AddTag(ctx context.Context, payload models.AddTagDTO) (int64, error)
	RemoveTag(ctx context.Context, payload models.RemoveTagDTO) (int64, error)
	SetSolver(ctx context.Context, payload models.SetSolverDTO) (int64, error)
	SetIsActive(ctx context.Context, taskId int64) error
	DeleteTask(ctx context.Context, taskId int64) (int64, error)
}

type ServerAPI struct {
	actionv1.UnimplementedActionServer
	commentv1.UnimplementedCommentServer
	progressbarv1.UnimplementedProgressBarServer
	statusv1.UnimplementedStatusServer
	tagv1.UnimplementedTagServer
	taskv1.UnimplementedTaskServer

	service  Service
	logger   zerolog.Logger
	validate *validator.Validate
}

func RegisterGRPCServer(
	grpcServer *grpc.Server,
	validate *validator.Validate,
	logger zerolog.Logger,
	service Service) {
	server := &ServerAPI{
		validate: validate,
		logger:   logger,
		service:  service,
	}

	actionv1.RegisterActionServer(grpcServer, server)
	commentv1.RegisterCommentServer(grpcServer, server)
	statusv1.RegisterStatusServer(grpcServer, server)
	tagv1.RegisterTagServer(grpcServer, server)
	taskv1.RegisterTaskServer(grpcServer, server)
	progressbarv1.RegisterProgressBarServer(grpcServer, server)
}

func (s *ServerAPI) ValidateStruct(data interface{}) error {
	if err := s.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}
