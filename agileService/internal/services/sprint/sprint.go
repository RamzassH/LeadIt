package sprint

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/rs/zerolog"
)

type Sprint struct {
	logger       zerolog.Logger
	saver        SprintSaver
	provider     SprintProvider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type SprintSaver interface {
	Save(ctx context.Context, payload models.CreateSprintDTO) (int64, error)
}

type SprintProvider interface {
	GetById(ctx context.Context, id int64) (*models.SprintDTO, error)
	GetManyByProjectId(ctx context.Context, projectId int64) ([]*models.SprintDTO, error)
	Update(ctx context.Context, payload models.UpdateSprintDTO) (*models.SprintDTO, error)
	Delete(ctx context.Context, id int64) (int64, error)
	SetIsActive(ctx context.Context, id int64) error
}

func NewSprint(logger zerolog.Logger, saver SprintSaver, provider SprintProvider, redis redisStorage.RedisStore, kafka *kafka.Producer) *Sprint {
	return &Sprint{
		logger:       logger,
		saver:        saver,
		provider:     provider,
		redisStorage: redis,
		kafka:        kafka,
	}
}

func (s *Sprint) CreateSprint(ctx context.Context, payload models.CreateSprintDTO) (int64, error) {
	const op = "Sprint.CreateSprint"
	s.logger.Info().Str("name", payload.Name).Msg("Creating sprint")
	id, err := s.saver.Save(ctx, payload)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to create sprint")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	s.logger.Info().Int64("sprint_id", id).Msg("Sprint created")
	return id, nil
}

func (s *Sprint) GetSprint(ctx context.Context, id int64) (*models.SprintDTO, error) {
	const op = "Sprint.GetSprint"
	s.logger.Debug().Int64("sprint_id", id).Msg("Fetching sprint")
	sprint, err := s.provider.GetById(ctx, id)
	if err != nil {
		s.logger.Error().Err(err).Msg("Sprint not found")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	s.logger.Info().Str("sprint_name", sprint.Name).Msg("Sprint fetched")
	return sprint, nil
}

func (s *Sprint) GetSprintsByProject(ctx context.Context, projectId int64) ([]*models.SprintDTO, error) {
	const op = "Sprint.GetSprintsByProject"
	s.logger.Debug().Int64("project_id", projectId).Msg("Fetching sprints")
	sprints, err := s.provider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch sprints")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	s.logger.Info().Int("sprint_count", len(sprints)).Msg("Sprints loaded")
	return sprints, nil
}

func (s *Sprint) UpdateSprint(ctx context.Context, payload models.UpdateSprintDTO) (*models.SprintDTO, error) {
	const op = "Sprint.UpdateSprint"
	s.logger.Info().Int64("sprint_id", payload.ID).Msg("Updating sprint")
	updated, err := s.provider.Update(ctx, payload)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to update sprint")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	s.logger.Info().Str("new_name", updated.Name).Msg("Sprint updated")
	return updated, nil
}

func (s *Sprint) DeleteSprint(ctx context.Context, id int64) (int64, error) {
	const op = "Sprint.DeleteSprint"
	s.logger.Warn().Int64("sprint_id", id).Msg("Deleting sprint")
	rows, err := s.provider.Delete(ctx, id)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to delete sprint")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	s.logger.Info().Int64("rows_deleted", rows).Msg("Sprint deleted")
	return rows, nil
}

func (s *Sprint) SetSprintIsActive(ctx context.Context, id int64) error {
	const op = "Sprint.SetSprintIsActive"
	s.logger.Debug().Int64("sprint_id", id).Msg("Toggling is_active")
	if err := s.provider.SetIsActive(ctx, id); err != nil {
		s.logger.Error().Err(err).Msg("Failed to toggle is_active")
		return fmt.Errorf("%s: %w", op, err)
	}
	s.logger.Info().Int64("sprint_id", id).Msg("is_active toggled")
	return nil
}
