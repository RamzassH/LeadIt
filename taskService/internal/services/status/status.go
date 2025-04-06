package status

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Status struct {
	logger         zerolog.Logger
	statusSaver    Saver
	statusProvider Provider
	redisStorage   redisStorage.RedisStore
	kafka          *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateStatusDTO) (int64, error)
}
type Provider interface {
	GetByProgressBarId(ctx context.Context, progressBarId int64) (status *models.StatusDTO, err error)
	GetManyByProjectId(ctx context.Context, projectId int64) (statusList []*models.StatusDTO, err error)
	Update(ctx context.Context, payload models.UpdateStatusDTO) (progressBar *models.StatusDTO, err error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	statusSaver Saver,
	statusProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *Status {
	return &Status{
		logger:         logger,
		statusSaver:    statusSaver,
		statusProvider: statusProvider,
		redisStorage:   redisStorage,
		kafka:          kafka,
	}
}

func (s *Status) CreateStatus(ctx context.Context, payload models.CreateStatusDTO) (int64, error) {
	const op = "Status.CreateStatus"
	logger := s.logger.With().
		Str("operation", op).
		Int64("progress_bar_id", payload.ProgressBarID).
		Logger()

	logger.Info().Str("status_name", payload.Name).Msg("Creating status")

	statusId, err := s.statusSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Str("status_name", payload.Name).
			Msg("Failed to create status")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("status_id", statusId).
		Msg("Status created")
	return statusId, nil
}

func (s *Status) GetStatusForProgressBar(ctx context.Context, progressBarId int64) (*models.StatusDTO, error) {
	const op = "Status.GetStatusForProgressBar"
	logger := s.logger.With().
		Str("operation", op).
		Int64("progress_bar_id", progressBarId).
		Logger()

	logger.Debug().Msg("Fetching status for progress bar")

	status, err := s.statusProvider.GetByProgressBarId(ctx, progressBarId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Progress bar status lookup failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug().
		Int64("status_id", status.ID).
		Msg("Status found")
	return status, nil
}

func (s *Status) GetStatusesForProject(ctx context.Context, projectId int64) ([]*models.StatusDTO, error) {
	const op = "Status.GetStatusesForProject"
	logger := s.logger.With().
		Str("operation", op).
		Int64("project_id", projectId).
		Logger()

	logger.Info().Msg("Fetching project statuses")

	statusList, err := s.statusProvider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to retrieve status list")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int("status_count", len(statusList)).
		Msg("Project statuses retrieved")
	return statusList, nil
}

func (s *Status) UpdateStatus(ctx context.Context, payload models.UpdateStatusDTO) (*models.StatusDTO, error) {
	const op = "Status.UpdateStatus"
	logger := s.logger.With().
		Str("operation", op).
		Int64("status_id", payload.ID).
		Logger()

	logger.Info().Msg("Updating status metadata")

	updatedStatus, err := s.statusProvider.Update(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Interface("new_values", payload).
			Msg("Status update failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Str("new_name", updatedStatus.Name).
		Msg("Status updated")
	return updatedStatus, nil
}

func (s *Status) DeleteStatus(ctx context.Context, statusId int64) (int64, error) {
	const op = "Status.DeleteStatus"
	logger := s.logger.With().
		Str("operation", op).
		Int64("status_id", statusId).
		Logger()

	logger.Warn().Msg("Initiating status deletion")

	rowsAffected, err := s.statusProvider.Delete(ctx, statusId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Critical: status deletion failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Warn().
		Int64("rows_affected", rowsAffected).
		Msg("Status permanently deleted")
	return rowsAffected, nil
}
