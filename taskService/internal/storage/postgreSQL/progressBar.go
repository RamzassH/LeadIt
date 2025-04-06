package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/RamzassH/LeadIt/taskService/internal/storage"
	"github.com/jmoiron/sqlx"
)

type ProgressBarStorage struct {
	db *sqlx.DB
}

func NewProgressBarStorage(db *sqlx.DB) (*ProgressBarStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &ProgressBarStorage{db: db}, nil
}

func (s *ProgressBarStorage) Save(ctx context.Context, payload models.CreateProgressBarDTO) (progressBarId int64, err error) {
	const op = "ProgressBarStorage.Save"
	query := `INSERT INTO progressbars (name, project_id) VALUES ($1, $2) RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.ProjectID).Scan(&progressBarId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return progressBarId, nil
}

func (s *ProgressBarStorage) GetManyByProjectId(ctx context.Context, projectId int64) (progressBar []*models.ProgressBarDTO, err error) {
	const op = "ProgressBarStorage.GetManyByProjectId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM progressbars WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.ProgressBarDTO
		if err := rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		progressBar = append(progressBar, &item)
	}
	return progressBar, nil
}

func (s *ProgressBarStorage) Update(ctx context.Context, payload models.UpdateProgressBarDTO) (*models.ProgressBarDTO, error) {
	const op = "ProgressBarStorage.Update"
	query := `UPDATE progressbars 
				SET 
                    name = COALESCE($1, name)
                WHERE id = $2 RETURNING *;`

	var updated models.ProgressBarDTO

	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.ID)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *ProgressBarStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "ProgressBarStorage.Delete"
	rowsAffected, err := storage.Delete(ctx, s.db, "progress_bars", id)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
