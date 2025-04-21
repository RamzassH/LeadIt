package taskBoard

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/rs/zerolog"
)

type TaskBoard struct {
	logger       zerolog.Logger
	saver        TaskBoardSaver
	provider     TaskBoardProvider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type TaskBoardSaver interface {
	Attach(ctx context.Context, payload models.AttachTaskToBoardDTO) (bool, error)
}

type TaskBoardProvider interface {
	Move(ctx context.Context, payload models.MoveTaskDTO) (bool, error)
	Detach(ctx context.Context, payload models.TaskBoardDTO) (int64, error)
}

func NewTaskBoard(logger zerolog.Logger, saver TaskBoardSaver, provider TaskBoardProvider, redis redisStorage.RedisStore, kafka *kafka.Producer) *TaskBoard {
	return &TaskBoard{
		logger:       logger,
		saver:        saver,
		provider:     provider,
		redisStorage: redis,
		kafka:        kafka,
	}
}

func (tb *TaskBoard) AttachTaskToBoard(ctx context.Context, payload models.AttachTaskToBoardDTO) (bool, error) {
	const op = "TaskBoard.AttachTaskToBoard"
	tb.logger.Info().Int64("task_id", payload.TaskID).Int64("board_id", payload.BoardID).Msg("Attaching task to board")
	success, err := tb.saver.Attach(ctx, payload)
	if err != nil {
		tb.logger.Error().Err(err).Msg("Failed to attach task")
		return false, fmt.Errorf("%s: %w", op, err)
	}
	tb.logger.Info().Bool("status", success).Msg("Task attached to board")
	return success, nil
}

func (tb *TaskBoard) MoveTask(ctx context.Context, payload models.MoveTaskDTO) (bool, error) {
	const op = "TaskBoard.MoveTask"
	tb.logger.Info().Int64("task_id", payload.TaskID).Int64("from", payload.FromColumnID).Int64("to", payload.ToColumnID).Msg("Moving task between columns")
	success, err := tb.provider.Move(ctx, payload)
	if err != nil {
		tb.logger.Error().Err(err).Msg("Move failed")
		return false, fmt.Errorf("%s: %w", op, err)
	}
	tb.logger.Info().Bool("success", success).Msg("Task moved")
	return success, nil
}

func (tb *TaskBoard) DetachTaskFromBoard(ctx context.Context, payload models.TaskBoardDTO) (int64, error) {
	const op = "TaskBoard.DetachTaskFromBoard"
	tb.logger.Warn().Int64("task_id", payload.TaskID).Int64("board_id", payload.BoardID).Msg("Detaching task")
	rows, err := tb.provider.Detach(ctx, payload)
	if err != nil {
		tb.logger.Error().Err(err).Msg("Detach failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	tb.logger.Info().Int64("rows", rows).Msg("Task detached")
	return rows, nil
}
