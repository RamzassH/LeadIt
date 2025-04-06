package comment

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Comment struct {
	logger       zerolog.Logger
	saver        Saver
	provider     Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateCommentDTO) (int64, error)
}
type Provider interface {
	GetManyByTaskId(ctx context.Context, id int64) (comment []*models.CommentDTO, err error)
	Update(ctx context.Context, payload models.UpdateCommentDTO) (comment *models.CommentDTO, err error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	commentSaver Saver,
	commentProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *Comment {
	return &Comment{
		logger:       logger,
		saver:        commentSaver,
		provider:     commentProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (c *Comment) CreateComment(ctx context.Context, payload models.CreateCommentDTO) (int64, error) {
	const op = "Comment.AddComment"
	logger := c.logger.With().
		Str("operation", op).
		Int64("task_id", payload.TaskID).
		Int64("author_id", payload.UserID).
		Logger()

	logger.Info().Msg("Posting new comment")

	commentId, err := c.saver.Save(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to post comment")
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("comment_id", commentId).
		Msg("Comment published")
	return commentId, nil
}

func (c *Comment) GetCommentsForTask(ctx context.Context, taskId int64) ([]*models.CommentDTO, error) {
	const op = "Comment.GetCommentsForTask"
	logger := c.logger.With().
		Str("operation", op).
		Int64("task_id", taskId).
		Logger()

	logger.Debug().Msg("Fetching task comments")

	comments, err := c.provider.GetManyByTaskId(ctx, taskId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Comment retrieval failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int("comment_count", len(comments)).
		Msg("Comments loaded")
	return comments, nil
}

func (c *Comment) UpdateComment(ctx context.Context, payload models.UpdateCommentDTO) (*models.CommentDTO, error) {
	const op = "Comment.UpdateComment"
	logger := c.logger.With().
		Str("operation", op).
		Int64("comment_id", payload.ID).
		Logger()

	logger.Info().Msg("Modifying comment content")

	updatedComment, err := c.provider.Update(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Int64("editor_id", payload.UserID).
			Msg("Comment update rejected")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Time("edited_at", updatedComment.Date).
		Msg("Comment modified")
	return updatedComment, nil
}

func (c *Comment) DeleteComment(ctx context.Context, commentId int64) (int64, error) {
	const op = "Comment.DeleteComment"
	logger := c.logger.With().
		Str("operation", op).
		Int64("comment_id", commentId).
		Logger()

	logger.Warn().Msg("Processing comment deletion request")

	rowsAffected, err := c.provider.Delete(ctx, commentId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Comment deletion failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Warn().
		Int64("rows_affected", rowsAffected).
		Msg("Comment permanently removed")
	return rowsAffected, nil
}
