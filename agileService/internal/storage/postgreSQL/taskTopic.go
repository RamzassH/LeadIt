package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type TaskTopicStorage struct {
	db *sqlx.DB
}

func NewTaskTopicStorage(db *sqlx.DB) *TaskTopicStorage {
	return &TaskTopicStorage{db: db}
}

func (s *TaskTopicStorage) Attach(ctx context.Context, payload models.TaskTopicDTO) (bool, error) {
	const op = "TaskTopicStorage.Attach"
	query := `INSERT INTO task_topic (task_id, topic_id) VALUES ($1, $2);`
	_, err := s.db.ExecContext(ctx, query, payload.TaskID, payload.TopicID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return true, nil
}

func (s *TaskTopicStorage) Detach(ctx context.Context, payload models.TaskTopicDTO) (int64, error) {
	const op = "TaskTopicStorage.Detach"
	query := `DELETE FROM task_topic WHERE task_id = $1 AND topic_id = $2;`
	res, err := s.db.ExecContext(ctx, query, payload.TaskID, payload.TopicID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return res.RowsAffected()
}
