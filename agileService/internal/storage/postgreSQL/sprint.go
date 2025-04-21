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

type SprintStorage struct {
	db *sqlx.DB
}

func NewSprintStorage(db *sqlx.DB) (*SprintStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &SprintStorage{db: db}, nil
}

func (s *SprintStorage) Save(ctx context.Context, payload models.CreateSprintDTO) (sprintId int64, err error) {
	const op = "SprintStorage.Save"

	query := `
	INSERT INTO sprints (project_id, name, description, start_date, end_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.ProjectID, payload.Name, payload.Description, payload.StartDate, payload.EndDate).Scan(&sprintId)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return sprintId, nil
}
func (s *SprintStorage) GetById(ctx context.Context, id int64) (sprint *models.SprintDTO, err error) {
	const op = "SprintStorage.GetById"

	err = storage.GetById(ctx, s.db, "sprints", id, &sprint)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return sprint, nil
}

func (s *SprintStorage) GetManyByProjectId(ctx context.Context, projectId int64) (sprints []*models.SprintDTO, err error) {
	const op = "SprintStorage.GetManyByProjectId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM sprints WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var sprint models.SprintDTO
		if err := rows.StructScan(&sprint); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		sprints = append(sprints, &sprint)
	}

	return sprints, nil
}
func (s *SprintStorage) Update(ctx context.Context, payload models.UpdateSprintDTO) (updated *models.SprintDTO, err error) {
	const op = "SprintStorage.Update"

	query := `UPDATE sprints 
				SET name = $1, description = $2, start_date = $3, end_date = $4, is_active = $5
				WHERE id = $6
				RETURNING *;`

	err = s.db.GetContext(ctx, &updated, query, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return updated, nil
}

func (s *SprintStorage) SetIsActive(ctx context.Context, id int64) error {
	const op = "SprintStorage.SetIsActive"

	query := `UPDATE sprints 
				SET
				is_active NOT is_active 
				WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (s *SprintStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "SprintStorage.Delete"

	rowsAffected, err := storage.Delete(ctx, s.db, "sprints", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
