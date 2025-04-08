package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/RamzassH/LeadIt/taskService/internal/storage"
	"github.com/jmoiron/sqlx"
)

type StatusStorage struct {
	db *sqlx.DB
}

func NewStatusStorage(db *sqlx.DB) (*StatusStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &StatusStorage{db: db}, nil
}

func (s *StatusStorage) Save(ctx context.Context, payload models.CreateStatusDTO) (statusId int64, err error) {
	const op = "StatusStorage.Save"
	query := `INSERT INTO statuses (name, order, progress_bar_id, notify_roles_ids) VALUES ($1, $2, $3, $4) RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.Order, payload.ProgressBarID, payload.NotifyRolesIDS).Scan(&statusId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return statusId, nil
}

func (s *StatusStorage) GetByProgressBarId(ctx context.Context, progressBarId int64) (status *models.StatusDTO, err error) {
	const op = "StatusStorage.GetByProgressBarId"

	err = storage.GetById(ctx, s.db, "statuses", progressBarId, &status)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return status, nil
}

func (s *StatusStorage) GetManyByProjectId(ctx context.Context, projectId int64) (statusList []*models.StatusDTO, err error) {
	const op = "StatusStorage.GetManyByProjectId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM statuses WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.StatusDTO
		if err := rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		statusList = append(statusList, &item)
	}
	return statusList, nil
}

func (s *StatusStorage) Update(ctx context.Context, payload models.UpdateStatusDTO) (*models.StatusDTO, error) {
	const op = "StatusStorage.Update"
	query := `UPDATE statuses SET 
                   name = COALESCE($1, name),
                   order = COALESCE($2, order),
                   notify_roles_ids = COALESCE($3, notify_roles_ids)
                   WHERE id = $4 RETURNING *;`
	var updated models.StatusDTO
	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.Order, payload.NotifyRolesIDS, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &updated, nil
}

func (s *StatusStorage) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "StatusStorage.Delete"

	rowsAffected, err = storage.Delete(ctx, s.db, "statuses", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
