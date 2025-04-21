package postgreSQL

import (
	"context"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/RamzassH/LeadIt/agileService/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type BoardColumnStorage struct {
	db *sqlx.DB
}

func NewBoardColumnStorage(db *sqlx.DB) (*BoardColumnStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	return &BoardColumnStorage{db: db}, nil
}

func (s *BoardColumnStorage) Save(ctx context.Context, payload models.CreateBoardColumnDTO) (int64, error) {
	const op = "BoardColumnStorage.Save"

	query := `
		INSERT INTO board_columns (board_id, name, order_num)
		VALUES ($1, $2, $3)
		RETURNING id;`

	var id int64
	err := s.db.QueryRowContext(ctx, query, payload.BoardID, payload.Name, payload.Order).Scan(&id)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code.Name() == "unique_violation" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *BoardColumnStorage) GetById(ctx context.Context, id int64) (*models.BoardColumnDTO, error) {
	const op = "BoardColumnStorage.GetById"

	var column models.BoardColumnDTO
	err := storage.GetById(ctx, s.db, "board_columns", id, &column)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &column, nil
}

func (s *BoardColumnStorage) GetManyByBoardId(ctx context.Context, boardId int64) ([]*models.BoardColumnDTO, error) {
	const op = "BoardColumnStorage.GetManyByBoardId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM board_columns WHERE board_id = $1 ORDER BY order_num`, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var columns []*models.BoardColumnDTO
	for rows.Next() {
		var column models.BoardColumnDTO
		if err := rows.StructScan(&column); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

func (s *BoardColumnStorage) Update(ctx context.Context, payload models.UpdateBoardColumnDTO) (*models.BoardColumnDTO, error) {
	const op = "BoardColumnStorage.Update"

	query := `
		UPDATE board_columns
		SET name = COALESCE($1, name), order_num = COALESCE($2, order_num)
		WHERE id = $3
		RETURNING *;`

	var updated models.BoardColumnDTO
	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.Order, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *BoardColumnStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "BoardColumnStorage.Delete"

	rowsAffected, err := storage.Delete(ctx, s.db, "board_columns", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
