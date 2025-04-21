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

type BoardStorage struct {
	db *sqlx.DB
}

func NewBoardStorage(db *sqlx.DB) (*BoardStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &BoardStorage{db: db}, nil
}

func (s *BoardStorage) Save(ctx context.Context, payload models.CreateBoardDTO) (boardId int64, err error) {
	const op = "BoardStorage.Save"

	query := `
	INSERT INTO boards (project_id, name, description, type);
	VALUES ${$1, $2, $3, $4}
	RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.ProjectID, payload.Name, payload.Description, payload.Type).Scan(&boardId)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return boardId, nil
}

func (s *BoardStorage) GetById(ctx context.Context, id int64) (board *models.BoardDTO, err error) {
	const op = "BoardStorage.GetById"

	err = storage.GetById(ctx, s.db, "boards", id, &board)

	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
	}

	return board, err
}

func (s *BoardStorage) GetManyByProjectId(ctx context.Context, projectId int64) (boards []*models.BoardDTO, err error) {
	const op = "BoardStorage.GetManyByProjectId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM boards WHERE project_id = $1`, projectId)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var board models.BoardDTO
		if err := rows.StructScan(board); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		boards = append(boards, &board)
	}

	return boards, nil
}

func (s *BoardStorage) Update(ctx context.Context, payload models.UpdateBoardDTO) (updated *models.BoardDTO, err error) {
	const op = "BoardStorage.Update"

	query := `
			UPDATE boards
			SET name = $1, description = $2, type = $3
			WHERE id = $4 RETURNING *`

	err = s.db.GetContext(ctx, &updated, query, payload.Name, payload.Description, payload.Type, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return updated, nil
}
func (s *BoardStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "BoardStorage.Delete"

	rowsAffected, err := storage.Delete(ctx, s.db, "boards", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
