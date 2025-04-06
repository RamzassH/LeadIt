package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/RamzassH/LeadIt/taskService/internal/storage"
	"github.com/jmoiron/sqlx"
)

type TagStorage struct {
	db *sqlx.DB
}

func NewTagStorage(db *sqlx.DB) (*TagStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &TagStorage{db: db}, nil
}

func (s *TagStorage) Save(ctx context.Context, payload models.CreateTagDTO) (tagId int64, err error) {
	const op = "TagStorage.Save"
	query := `INSERT INTO tags (name, tag_color, project_id) VALUES ($1, $2, $3) RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.TagColor, payload.ProjectID).Scan(&tagId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return tagId, nil
}

func (s *TagStorage) GetManyByProjectId(ctx context.Context, projectId int64) (tagList []*models.TagDTO, err error) {
	const op = "TagStorage.GetManyByProjectId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM tags WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.TagDTO
		if err := rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tagList = append(tagList, &item)
	}
	return tagList, nil
}

func (s *TagStorage) Update(ctx context.Context, payload models.UpdateTagDTO) (*models.TagDTO, error) {
	const op = "TagStorage.Update"
	query := `UPDATE tags 
				SET 
                name = COALESCE($1, name), 
                tag_color = COALESCE($2, tag_color)
              WHERE id = $3 RETURNING *;`
	var updated models.TagDTO
	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.TagColor, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *TagStorage) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "TagStorage.Delete"

	rowsAffected, err = storage.Delete(ctx, s.db, "tags", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
