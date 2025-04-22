package tag

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Tag struct {
	logger       zerolog.Logger
	tagSaver     Saver
	tagProvider  Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateTagDTO) (int64, error)
}
type Provider interface {
	GetManyByProjectId(ctx context.Context, projectId int64) (tags []*models.TagDTO, err error)
	Update(ctx context.Context, payload models.UpdateTagDTO) (tag *models.TagDTO, err error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	tagSaver Saver,
	tagProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *Tag {
	return &Tag{
		logger:       logger,
		tagSaver:     tagSaver,
		tagProvider:  tagProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (t *Tag) CreateTag(ctx context.Context, payload models.CreateTagDTO) (int64, error) {
	const op = "Tag.CreateTag"
	logger := t.logger.With().Str("operation", op).Logger()
	logger.Info().Int64("project_id", payload.ProjectID).Msg("Starting tag creation")

	tagId, err := t.tagSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Int64("project_id", payload.ProjectID).
			Str("tag_name", payload.Name).
			Msg("Failed to create tag")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("tag_id", tagId).
		Str("tag_name", payload.Name).
		Msg("Tag created successfully")
	return tagId, nil
}

func (t *Tag) GetTagsForProject(ctx context.Context, projectId int64) ([]*models.TagDTO, error) {
	const op = "Tag.GetTagsForProject"
	logger := t.logger.With().
		Str("operation", op).
		Int64("project_id", projectId).
		Logger()
	logger.Info().Msg("Fetching project tags")

	tags, err := t.tagProvider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to fetch tags")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int("tags_count", len(tags)).
		Msg("Project tags retrieved")
	return tags, nil
}

func (t *Tag) UpdateTag(ctx context.Context, payload models.UpdateTagDTO) (*models.TagDTO, error) {
	const op = "Tag.UpdateTag"
	logger := t.logger.With().
		Str("operation", op).
		Int64("tag_id", payload.ID).
		Logger()
	logger.Info().Msg("Updating tag")

	updatedTag, err := t.tagProvider.Update(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Tag update failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Str("new_name", updatedTag.Name).
		Msg("Tag updated")
	return updatedTag, nil
}

func (t *Tag) DeleteTag(ctx context.Context, tagId int64) (int64, error) {
	const op = "Tag.DeleteTag"
	logger := t.logger.With().
		Str("operation", op).
		Int64("tag_id", tagId).
		Logger()
	logger.Info().Msg("Starting tag deletion")

	rowsAffected, err := t.tagProvider.Delete(ctx, tagId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Tag deletion failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("rows_affected", rowsAffected).
		Msg("Tag deleted")
	return rowsAffected, nil
}
