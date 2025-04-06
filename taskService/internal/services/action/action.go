package action

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Action struct {
	logger       zerolog.Logger
	saver        Saver
	provider     Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateActionDTO) (int64, error)
}
type Provider interface {
	GetManyByTaskId(ctx context.Context, id int64) (actions []*models.ActionDTO, err error)
}

func New(
	logger zerolog.Logger,
	actionSaver Saver,
	actionProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *Action {
	return &Action{
		logger:       logger,
		saver:        actionSaver,
		provider:     actionProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (a *Action) CreateAction(ctx context.Context, payload models.CreateActionDTO) (int64, error) {
	const op = "Action.AddAction"
	logger := a.logger.With().
		Str("operation", op).
		Int64("task_id", payload.TaskID).
		Str("action_type", payload.ActionType).
		Logger()

	logger.Info().Msg("Registering new action")

	actionID, err := a.saver.Save(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Interface("payload", payload).
			Msg("Action registration failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("action_id", actionID).
		Msg("Action recorded")
	return actionID, nil
}

func (a *Action) GetActionsForTask(ctx context.Context, taskId int64) ([]*models.ActionDTO, error) {
	const op = "Action.GetActionsForTask"
	logger := a.logger.With().
		Str("operation", op).
		Int64("task_id", taskId).
		Logger()

	logger.Debug().Msg("Auditing task actions")

	actions, err := a.
		provider.GetManyByTaskId(ctx, taskId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Action audit failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int("action_count", len(actions)).
		Str("first_action_type", safeFirstActionType(actions)).
		Msg("Action audit complete")
	return actions, nil
}

func safeFirstActionType(actions []*models.ActionDTO) string {
	if len(actions) == 0 {
		return "none"
	}
	return actions[0].ActionType
}
