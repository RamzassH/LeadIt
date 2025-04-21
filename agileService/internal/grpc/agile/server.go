package agile

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	boardV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/board"
	boardColumnV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/board_column"
	sprintV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/sprint"
	taskBoardV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/task_board"
	taskTopicV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/task_topic"
	topicV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/topic"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Service interface {
	Board
	BoardColumn
	Sprint
	TaskBoard
	TaskTopic
	Topic
}

type Board interface {
	CreateBoard(ctx context.Context, payload models.CreateBoardDTO) (int64, error)
	GetBoard(ctx context.Context, id int64) (*models.BoardDTO, error)
	GetBoardsByProject(ctx context.Context, projectId int64) ([]*models.BoardDTO, error)
	UpdateBoard(ctx context.Context, payload models.UpdateBoardDTO) (*models.BoardDTO, error)
	DeleteBoard(ctx context.Context, id int64) (int64, error)
}

type BoardColumn interface {
	CreateBoardColumn(ctx context.Context, payload models.CreateBoardColumnDTO) (int64, error)
	GetBoardColumnsByBoard(ctx context.Context, boardId int64) ([]*models.BoardColumnDTO, error)
	GetBoardColumn(ctx context.Context, id int64) (*models.BoardColumnDTO, error)
	UpdateBoardColumn(ctx context.Context, payload models.UpdateBoardColumnDTO) (*models.BoardColumnDTO, error)
	DeleteBoardColumn(ctx context.Context, id int64) (int64, error)
}

type Sprint interface {
	CreateSprint(ctx context.Context, payload models.CreateSprintDTO) (int64, error)
	GetSprint(ctx context.Context, id int64) (*models.SprintDTO, error)
	GetSprintsByProject(ctx context.Context, projectId int64) ([]*models.SprintDTO, error)
	UpdateSprint(ctx context.Context, payload models.UpdateSprintDTO) (*models.SprintDTO, error)
	DeleteSprint(ctx context.Context, id int64) (int64, error)
	SetSprintIsActive(ctx context.Context, sprintId int64) error
}

type TaskBoard interface {
	AttachTaskToBoard(ctx context.Context, payload models.AttachTaskToBoardDTO) (bool, error)
	DetachTaskFromBoard(ctx context.Context, payload models.TaskBoardDTO) (int64, error)
	MoveTask(ctx context.Context, payload models.MoveTaskDTO) (bool, error)
}

type TaskTopic interface {
	AttachTopicToTask(ctx context.Context, payload models.TaskTopicDTO) (bool, error)
	DetachTopicFromTask(ctx context.Context, payload models.TaskTopicDTO) (int64, error)
}

type Topic interface {
	CreateTopic(ctx context.Context, payload models.CreateTopicDTO) (int64, error)
	GetTopic(ctx context.Context, id int64) (*models.TopicDTO, error)
	GetTopicsByProject(ctx context.Context, projectId int64) ([]*models.TopicDTO, error)
	UpdateTopic(ctx context.Context, payload models.UpdateTopicDTO) (*models.TopicDTO, error)
	DeleteTopic(ctx context.Context, id int64) (int64, error)
}

type ServerApi struct {
	boardV1.UnimplementedBoardServiceServer
	boardColumnV1.UnimplementedBoardColumnServiceServer
	sprintV1.UnimplementedSprintServiceServer
	taskBoardV1.UnimplementedTaskBoardServiceServer
	taskTopicV1.UnimplementedTaskTopicServiceServer
	topicV1.UnimplementedTopicServiceServer

	service  Service
	logger   zerolog.Logger
	validate *validator.Validate
}

func RegisterGRPCServer(
	grpcServer *grpc.Server,
	validate *validator.Validate,
	service Service,
	logger zerolog.Logger) {
	server := &ServerApi{
		validate: validate,
		logger:   logger,
		service:  service,
	}

	boardV1.RegisterBoardServiceServer(grpcServer, server)
	boardColumnV1.RegisterBoardColumnServiceServer(grpcServer, server)
	sprintV1.RegisterSprintServiceServer(grpcServer, server)
	taskBoardV1.RegisterTaskBoardServiceServer(grpcServer, server)
	taskTopicV1.RegisterTaskTopicServiceServer(grpcServer, server)
	topicV1.RegisterTopicServiceServer(grpcServer, server)
}

func (s *ServerApi) ValidateStruct(data interface{}) error {
	if err := s.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}
