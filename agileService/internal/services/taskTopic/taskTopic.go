package taskTopic

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/rs/zerolog"
)

type TaskTopic struct {
	logger       zerolog.Logger
	saver        TaskTopicSaver
	provider     TaskTopicProvider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type TaskTopicSaver interface {
	Attach(ctx context.Context, payload models.TaskTopicDTO) (bool, error)
}

type TaskTopicProvider interface {
	Detach(ctx context.Context, payload models.TaskTopicDTO) (int64, error)
}

func (tt *TaskTopic) AttachTopicToTask(ctx context.Context, payload models.TaskTopicDTO) (bool, error) {
	const op = "TaskTopic.AttachTopicToTask"

	tt.logger.Info().
		Int64("topic_id", payload.TopicID).
		Int64("task_id", payload.TaskID).
		Msg("Attaching topic to task")

	success, err := tt.saver.Attach(ctx, payload)
	if err != nil {
		tt.logger.Error().Err(err).Msg("Failed to attach topic")
		return false, fmt.Errorf("%s: %w", op, err)
	}

	tt.logger.Info().Bool("status", success).Msg("Attached topic to task")
	return success, nil
}

func (tt *TaskTopic) DetachTopicFromTask(ctx context.Context, payload models.TaskTopicDTO) (int64, error) {
	const op = "TaskTopic.DetachTopicFromTask"

	tt.logger.Info().
		Int64("topic_id", payload.TopicID).
		Int64("task_id", payload.TaskID).
		Msg("Detaching topic from task")

	success, err := tt.provider.Detach(ctx, payload)

	if err != nil {
		tt.logger.Error().Err(err).Msg("Failed to detach topic")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return success, nil

}
