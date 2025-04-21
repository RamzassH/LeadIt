package boardColumn

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/rs/zerolog"
)

type BoardColumn struct {
	logger       zerolog.Logger
	saver        BoardColumnSaver
	provider     BoardColumnProvider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type BoardColumnSaver interface {
	Save(ctx context.Context, payload models.CreateBoardColumnDTO) (int64, error)
}

type BoardColumnProvider interface {
	GetById(ctx context.Context, id int64) (*models.BoardColumnDTO, error)
	GetManyByBoardId(ctx context.Context, boardId int64) ([]*models.BoardColumnDTO, error)
	Update(ctx context.Context, payload models.UpdateBoardColumnDTO) (*models.BoardColumnDTO, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

func NewBoardColumn(logger zerolog.Logger, saver BoardColumnSaver, provider BoardColumnProvider, redis redisStorage.RedisStore, kafka *kafka.Producer) *BoardColumn {
	return &BoardColumn{
		logger:       logger,
		saver:        saver,
		provider:     provider,
		redisStorage: redis,
		kafka:        kafka,
	}
}

func (b *BoardColumn) CreateBoardColumn(ctx context.Context, payload models.CreateBoardColumnDTO) (int64, error) {
	const op = "BoardColumn.CreateBoardColumn"
	b.logger.Info().Str("name", payload.Name).Msg("Creating board column")
	id, err := b.saver.Save(ctx, payload)
	if err != nil {
		b.logger.Error().Err(err).Msg("Failed to create board column")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	b.logger.Info().Int64("column_id", id).Msg("Board column created")
	return id, nil
}

func (b *BoardColumn) GetBoardColumn(ctx context.Context, id int64) (*models.BoardColumnDTO, error) {
	const op = "BoardColumn.GetBoardColumn"
	b.logger.Debug().Int64("column_id", id).Msg("Fetching board column")
	column, err := b.provider.GetById(ctx, id)
	if err != nil {
		b.logger.Error().Err(err).Msg("Board column not found")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	b.logger.Info().Str("column_name", column.Name).Msg("Board column fetched")
	return column, nil
}

func (b *BoardColumn) GetBoardColumnsByBoard(ctx context.Context, boardId int64) ([]*models.BoardColumnDTO, error) {
	const op = "BoardColumn.GetBoardColumnsByBoard"
	b.logger.Debug().Int64("board_id", boardId).Msg("Fetching columns by board")
	columns, err := b.provider.GetManyByBoardId(ctx, boardId)
	if err != nil {
		b.logger.Error().Err(err).Msg("Failed to fetch columns")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	b.logger.Info().Int("column_count", len(columns)).Msg("Board columns loaded")
	return columns, nil
}

func (b *BoardColumn) UpdateBoardColumn(ctx context.Context, payload models.UpdateBoardColumnDTO) (*models.BoardColumnDTO, error) {
	const op = "BoardColumn.UpdateBoardColumn"
	b.logger.Info().Int64("column_id", payload.ID).Msg("Updating board column")
	updated, err := b.provider.Update(ctx, payload)
	if err != nil {
		b.logger.Error().Err(err).Msg("Board column update failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	b.logger.Info().Str("new_name", updated.Name).Msg("Board column updated")
	return updated, nil
}

func (b *BoardColumn) DeleteBoardColumn(ctx context.Context, id int64) (int64, error) {
	const op = "BoardColumn.DeleteBoardColumn"
	b.logger.Warn().Int64("column_id", id).Msg("Deleting board column")
	rows, err := b.provider.Delete(ctx, id)
	if err != nil {
		b.logger.Error().Err(err).Msg("Failed to delete column")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	b.logger.Info().Int64("rows_deleted", rows).Msg("Board column deleted")
	return rows, nil
}
