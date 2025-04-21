package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type TaskBoardStorage struct {
	db *sqlx.DB
}

func NewTaskBoardStorage(db *sqlx.DB) *TaskBoardStorage {
	return &TaskBoardStorage{db: db}
}

func (s *TaskBoardStorage) Attach(ctx context.Context, payload models.AttachTaskToBoardDTO) (bool, error) {
	const op = "TaskBoardStorage.Attach"
	query := `INSERT INTO task_board (task_id, board_id, column_id) VALUES ($1, $2, $3);`
	_, err := s.db.ExecContext(ctx, query, payload.TaskID, payload.BoardID, payload.ColumnID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return true, nil
}

func (s *TaskBoardStorage) Move(ctx context.Context, payload models.MoveTaskDTO) (bool, error) {
	const op = "TaskBoardStorage.Move"
	query := `UPDATE task_board SET column_id = $1 WHERE task_id = $2 AND board_id = $3;`
	_, err := s.db.ExecContext(ctx, query, payload.ToColumnID, payload.TaskID, payload.BoardID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return true, nil
}

func (s *TaskBoardStorage) Detach(ctx context.Context, payload models.TaskBoardDTO) (int64, error) {
	const op = "TaskBoardStorage.Detach"
	query := `DELETE FROM task_board WHERE task_id = $1 AND board_id = $2;`
	res, err := s.db.ExecContext(ctx, query, payload.TaskID, payload.BoardID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return res.RowsAffected()
}
