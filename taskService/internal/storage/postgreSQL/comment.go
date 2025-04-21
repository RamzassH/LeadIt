package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/RamzassH/LeadIt/taskService/internal/storage"
	"github.com/jmoiron/sqlx"
)

type CommentStorage struct {
	db *sqlx.DB
}

func NewCommentStorage(db *sqlx.DB) (*CommentStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &CommentStorage{db: db}, nil
}

func (s *CommentStorage) Save(ctx context.Context, payload models.CreateCommentDTO) (commentId int64, err error) {
	const op = "CommentStorage.Save"
	query := `
	INSERT INTO comments (task_id, user_id, body)
	VALUES ($1, $2, $3)
 	RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.TaskID, payload.UserID, payload.Body).Scan(&commentId)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return commentId, nil
}

func (s *CommentStorage) GetManyByTaskId(ctx context.Context, id int64) (comments []*models.CommentDTO, err error) {
	const op = "CommentStorage.GetManyByTaskId"

	rows, err := s.db.QueryxContext(ctx,
		`SELECT * FROM comments WHERE task_id=$1`, id)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.CommentDTO
		if err := rows.StructScan(&comment); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (s *CommentStorage) Update(ctx context.Context, payload models.UpdateCommentDTO) (*models.CommentDTO, error) {
	const op = "CommentStorage.Update"
	query := `
	UPDATE comments 
	SET
	    body = COALESCE($1, body)
	WHERE id = $2 RETURNING *;`
	var updated models.CommentDTO

	err := s.db.GetContext(ctx, &updated, query, payload.Body, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *CommentStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "CommentStorage.Delete"

	rowsAffected, err := storage.Delete(ctx, s.db, "comments", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
