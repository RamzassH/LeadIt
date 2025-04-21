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

type TopicStorage struct {
	db *sqlx.DB
}

func NewTopicStorage(db *sqlx.DB) *TopicStorage {
	return &TopicStorage{db: db}
}

func (s *TopicStorage) Save(ctx context.Context, payload models.CreateTopicDTO) (int64, error) {
	const op = "TopicStorage.Save"
	query := `INSERT INTO topics (project_id, name, description) VALUES ($1, $2, $3) RETURNING id;`
	var id int64
	err := s.db.QueryRowContext(ctx, query, payload.ProjectID, payload.Name, payload.Description).Scan(&id)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code.Name() == "unique_violation" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *TopicStorage) GetById(ctx context.Context, id int64) (*models.TopicDTO, error) {
	const op = "TopicStorage.GetById"
	var topic models.TopicDTO
	err := storage.GetById(ctx, s.db, "topics", id, &topic)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &topic, nil
}

func (s *TopicStorage) GetManyByProjectId(ctx context.Context, projectId int64) ([]*models.TopicDTO, error) {
	const op = "TopicStorage.GetManyByProjectId"
	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM topics WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var topics []*models.TopicDTO
	for rows.Next() {
		var topic models.TopicDTO
		if err := rows.StructScan(&topic); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		topics = append(topics, &topic)
	}
	return topics, nil
}

func (s *TopicStorage) Update(ctx context.Context, payload models.UpdateTopicDTO) (*models.TopicDTO, error) {
	const op = "TopicStorage.Update"
	query := `UPDATE topics SET name = COALESCE($1, name), description = COALESCE($2, description) WHERE id = $3 RETURNING *;`
	var updated models.TopicDTO
	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.Description, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *TopicStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "TopicStorage.Delete"
	rowsAffected, err := storage.Delete(ctx, s.db, "topics", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
