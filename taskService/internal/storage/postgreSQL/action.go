package postgreSQL

import (
	"context"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/RamzassH/LeadIt/taskService/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ActionStorage struct {
	db *sqlx.DB
}

func NewActionStorage(db *sqlx.DB) (*ActionStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &ActionStorage{db: db}, nil
}

func (s *ActionStorage) Save(ctx context.Context, payload models.CreateActionDTO) (actionId int64, err error) {
	const op = "ActionStorage.Save"

	query := `
	INSERT INTO	actions (task_id, user_id, action_type, description)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.TaskID, payload.UserID, payload.ActionType, payload.Description).Scan(&actionId)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return actionId, nil
}

func (s *ActionStorage) GetManyByTaskId(ctx context.Context, id int64) (actions []*models.ActionDTO, err error) {
	const op = "ActionStorage.GetManyByTaskId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM actions WHERE task_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var action models.ActionDTO
		if err := rows.StructScan(&action); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		actions = append(actions, &action)
	}

	return actions, nil
}
