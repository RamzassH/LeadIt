package topic

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/rs/zerolog"
)

type Topic struct {
	logger       zerolog.Logger
	saver        TopicSaver
	provider     TopicProvider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type TopicSaver interface {
	Save(ctx context.Context, payload models.CreateTopicDTO) (int64, error)
}

type TopicProvider interface {
	GetById(ctx context.Context, id int64) (*models.TopicDTO, error)
	GetManyByProjectId(ctx context.Context, projectId int64) ([]*models.TopicDTO, error)
	Update(ctx context.Context, payload models.UpdateTopicDTO) (*models.TopicDTO, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

func NewTopic(logger zerolog.Logger, saver TopicSaver, provider TopicProvider, redis redisStorage.RedisStore, kafka *kafka.Producer) *Topic {
	return &Topic{
		logger:       logger,
		saver:        saver,
		provider:     provider,
		redisStorage: redis,
		kafka:        kafka,
	}
}

func (t *Topic) CreateTopic(ctx context.Context, payload models.CreateTopicDTO) (int64, error) {
	const op = "Topic.CreateTopic"
	t.logger.Info().Str("name", payload.Name).Msg("Creating topic")
	id, err := t.saver.Save(ctx, payload)
	if err != nil {
		t.logger.Error().Err(err).Msg("Failed to create topic")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	t.logger.Info().Int64("topic_id", id).Msg("Topic created")
	return id, nil
}

func (t *Topic) GetTopic(ctx context.Context, id int64) (*models.TopicDTO, error) {
	const op = "Topic.GetTopic"
	t.logger.Debug().Int64("topic_id", id).Msg("Fetching topic")
	topic, err := t.provider.GetById(ctx, id)
	if err != nil {
		t.logger.Error().Err(err).Msg("Topic not found")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	t.logger.Info().Str("topic_name", topic.Name).Msg("Topic fetched")
	return topic, nil
}

func (t *Topic) GetTopicsByProject(ctx context.Context, projectId int64) ([]*models.TopicDTO, error) {
	const op = "Topic.GetTopicsByProject"
	t.logger.Debug().Int64("project_id", projectId).Msg("Fetching topics by project")
	topics, err := t.provider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		t.logger.Error().Err(err).Msg("Failed to fetch topics")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	t.logger.Info().Int("topic_count", len(topics)).Msg("Topics loaded")
	return topics, nil
}

func (t *Topic) UpdateTopic(ctx context.Context, payload models.UpdateTopicDTO) (*models.TopicDTO, error) {
	const op = "Topic.UpdateTopic"
	t.logger.Info().Int64("topic_id", payload.ID).Msg("Updating topic")
	topic, err := t.provider.Update(ctx, payload)
	if err != nil {
		t.logger.Error().Err(err).Msg("Failed to update topic")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	t.logger.Info().Str("new_name", topic.Name).Msg("Topic updated")
	return topic, nil
}

func (t *Topic) DeleteTopic(ctx context.Context, id int64) (int64, error) {
	const op = "Topic.DeleteTopic"
	t.logger.Warn().Int64("topic_id", id).Msg("Deleting topic")
	rows, err := t.provider.Delete(ctx, id)
	if err != nil {
		t.logger.Error().Err(err).Msg("Failed to delete topic")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	t.logger.Info().Int64("rows_deleted", rows).Msg("Topic deleted")
	return rows, nil
}
