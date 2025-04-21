package board

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/rs/zerolog"
)

type Board struct {
	logger       zerolog.Logger
	saver        Saver
	provider     Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateBoardDTO) (int64, error)
}
type Provider interface {
	GetById(ctx context.Context, id int64) (*models.BoardDTO, error)
	GetManyByProjectId(ctx context.Context, projectId int64) ([]*models.BoardDTO, error)
	Update(ctx context.Context, payload models.UpdateBoardDTO) (*models.BoardDTO, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

func New(
	logger zerolog.Logger,
	boardSaver Saver,
	boardProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer) *Board {
	return &Board{
		logger:       logger,
		saver:        boardSaver,
		provider:     boardProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (b *Board) CreateBoard(ctx context.Context, payload models.CreateBoardDTO) (int64, error) {
	const op = "Board.CreateBoard"
	logger := b.logger.With().Str("operation", op).Logger()

	logger.Info().Str("name", payload.Name).Msg("Creating new board")
	id, err := b.saver.Save(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create board")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().Int64("board_id", id).Msg("Board created successfully")
	return id, nil
}

func (b *Board) GetBoard(ctx context.Context, id int64) (*models.BoardDTO, error) {
	const op = "Board.GetBoard"
	logger := b.logger.With().Str("operation", op).Logger()

	logger.Debug().Int64("board_id", id).Msg("Fetching board")
	board, err := b.provider.GetById(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("Board not found")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info().Str("board_name", board.Name).Msg("Board fetched")
	return board, nil
}

func (b *Board) GetBoardsByProject(ctx context.Context, projectId int64) ([]*models.BoardDTO, error) {
	const op = "Board.GetBoardsByProject"
	logger := b.logger.With().Str("operation", op).Int64("project_id", projectId).Logger()

	logger.Debug().Msg("Fetching all boards by project")
	boards, err := b.provider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get boards")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().Int("count", len(boards)).Msg("Boards loaded")
	return boards, nil
}

func (b *Board) UpdateBoard(ctx context.Context, payload models.UpdateBoardDTO) (*models.BoardDTO, error) {
	const op = "Board.UpdateBoard"
	logger := b.logger.With().Str("operation", op).Logger()

	logger.Info().Int64("board_id", payload.ID).Msg("Updating board")
	updatedBoard, err := b.provider.Update(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Msg("Board update failed")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().Str("new_name", updatedBoard.Name).Msg("Board updated")
	return updatedBoard, nil
}

func (b *Board) DeleteBoard(ctx context.Context, id int64) (int64, error) {
	const op = "Board.DeleteBoard"
	logger := b.logger.With().Str("operation", op).Logger()

	logger.Warn().Int64("board_id", id).Msg("Attempting to delete board")
	rows, err := b.provider.Delete(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("Delete board failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().Int64("rows_deleted", rows).Msg("Board deleted")
	return rows, nil
}
