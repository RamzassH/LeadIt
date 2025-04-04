package progressBar

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
)

type ProgressBar struct {
	logger       zerolog.Logger
	saver        Saver
	provider     Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateProgressBarDTO) (int64, error)
}
type Provider interface {
	GetManyByProjectId(ctx context.Context, id int64) (progressBar []*models.ProgressBarDTO, err error)
	Update(ctx context.Context, payload models.UpdateProgressBarDTO) (progressBar *models.ProgressBarDTO, err error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	progressBarSaver Saver,
	progressBarProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *ProgressBar {
	return &ProgressBar{
		logger:       logger,
		saver:        progressBarSaver,
		provider:     progressBarProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (pb *ProgressBar) CreateProgressBar(ctx context.Context, payload models.CreateProgressBarDTO) (int64, error) {
	const op = "ProgressBar.CreateProgressBar"
	logger := pb.logger.With().
		Str("operation", op).
		Int64("project_id", payload.ProjectID).
		Logger()

	logger.Info().Msg("Initializing progress bar creation")

	progressBarId, err := pb.saver.Save(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Str("progress_bar_name", payload.Name).
			Msg("Failed to initialize progress bar")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("progress_bar_id", progressBarId).
		Msg("Progress bar created")
	return progressBarId, nil
}

func (pb *ProgressBar) GetProgressBarForProject(ctx context.Context, projectId int64) (*models.ProgressBarDTO, error) {
	const op = "ProgressBar.GetProgressBarForProject"
	logger := pb.logger.With().
		Str("operation", op).
		Int64("project_id", projectId).
		Logger()

	logger.Debug().Msg("Looking up project progress bar")

	progressBarList, err := pb.provider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Progress bar not found")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("progress_bar_id", progressBar.ID).
		Msg("Progress bar retrieved")
	return progressBar, nil
}

func (pb *ProgressBar) UpdateProgressBar(ctx context.Context, payload models.UpdateProgressBarDTO) (*models.ProgressBarDTO, error) {
	const op = "ProgressBar.UpdateProgressBar"
	logger := pb.logger.With().
		Str("operation", op).
		Int64("progress_bar_id", payload.ID).
		Logger()

	logger.Info().Interface("update_data", payload).Msg("Starting update")

	updatedProgressBar, err := pb.provider.Update(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Progress bar update failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Msg("Progress bar updated")
	return updatedProgressBar, nil
}

func (pb *ProgressBar) DeleteProgressBar(ctx context.Context, progressBarId int64) (int64, error) {
	const op = "ProgressBar.DeleteProgressBar"
	logger := pb.logger.With().
		Str("operation", op).
		Int64("progress_bar_id", progressBarId).
		Logger()

	logger.Warn().Msg("Initiating progress bar deletion")

	rowsAffected, err := pb.provider.Delete(ctx, progressBarId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Critical: deletion failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Warn().
		Int64("rows_affected", rowsAffected).
		Msg("Progress bar destroyed")
	return rowsAffected, nil
}
