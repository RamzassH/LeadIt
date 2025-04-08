package postgreSQL

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/RamzassH/LeadIt/taskService/internal/storage"
	"github.com/jmoiron/sqlx"
)

type TaskStorage struct {
	db *sqlx.DB
}

func NewTaskStorage(db *sqlx.DB) (*TaskStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}
	return &TaskStorage{db: db}, nil
}
func (s *TaskStorage) Save(ctx context.Context, payload models.CreateTaskDTO) (taskId int64, err error) {
	const op = "TaskStorage.Save"
	query := `INSERT INTO tasks
    (name, description, project_id, setter_id, solver_id, allocated_time, progress_bar_id, current_status_id ) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.Description, payload.ProjectID, payload.SetterID,
		payload.SolverID, payload.AllocatedTime, payload.ProgressBarID, payload.CurrentStatusID).
		Scan(&taskId)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return taskId, nil
}

func (s *TaskStorage) GetById(ctx context.Context, taskId int64) (task *models.TaskDTO, err error) {
	const op = "TaskStorage.GetById"

	err = storage.GetById(ctx, s.db, "tasks", taskId, &task)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return task, nil
}

func (s *TaskStorage) GetManyByProjectId(ctx context.Context, projectId int64) (tasList []*models.TaskDTO, err error) {
	const op = "TaskStorage.GetManyByProjectId"

	rows, err := s.db.QueryxContext(ctx, `SELECT * FROM tasks WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.TaskDTO
		if err := rows.StructScan(&item); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasList = append(tasList, &item)
	}
	return tasList, nil
}

func (s *TaskStorage) Update(ctx context.Context, payload models.UpdateTaskDTO) (*models.TaskDTO, error) {
	const op = "TaskStorage.Update"
	query := `UPDATE tasks 
				SET 
                	name = COALESCE($1, name),
                	description = COALESCE($2, description),
                	solver_id = COALESCE($3, solver_id),
                	allocated_time = COALESCE($4, allocated_time),
                	last_start_time = COALESCE($5, last_start_time),
                	sprint_id = COALESCE($6, sprint_id),
                	wasted_time = COALESCE($7, wasted_time),
                	progress_bar_id = COALESCE($8, progress_bar_id),
                	current_status_id = COALESCE($9, current_status_id)
                 WHERE id = $10 RETURNING *;`
	var updated models.TaskDTO
	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.Description,
		payload.SolverID, payload.AllocatedTime,
		payload.SprintID, payload.WastedTime,
		payload.ProgressBarID, payload.CurrentStatusID, payload.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *TaskStorage) ChangeStatus(ctx context.Context, payload models.ChangeStatusDTO) (statusId int64, err error) {
	const op = "TaskStorage.ChangeStatus"

	query := `UPDATE tasks SET
				current_status_id = COALESCE($1, current_status_id)
				WHERE id = $2 RETURNING current_status_id;`

	err = s.db.GetContext(ctx, &statusId, query, payload.StatusID, payload.TaskID)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return statusId, nil
}

func (s *TaskStorage) AddTag(ctx context.Context, payload models.AddTagDTO) (tagId int64, err error) {
	const op = "TaskStorage.AddTag"

	query := `INSERT INTO task_tags 
			(id, tag_id) VALUES ($1, $2) 
			RETURNING tag_id;`

	err = s.db.QueryRowContext(ctx, query, payload.TaskID, payload.TagID).Scan(&tagId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return tagId, nil
}

func (s *TaskStorage) RemoveTag(ctx context.Context, payload models.RemoveTagDTO) (int64, error) {
	const op = "TaskStorage.RemoveTag"

	query := `DELETE FROM task_tags WHERE task_id = $1 AND tag_id = $2`

	result, err := s.db.ExecContext(ctx, query, payload.TaskID, payload.TagID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s (rowsAffected): %w", op, err)
	}

	return rowsAffected, nil
}

func (s *TaskStorage) SetSolver(ctx context.Context, payload models.SetSolverDTO) (solverId int64, err error) {
	const op = "TaskStorage.SetSolver"

	query := `UPDATE tasks SET solver_id = $1 WHERE id = $2 RETURNING solver_id;`

	err = s.db.GetContext(ctx, &solverId, query, payload.TaskID, payload.SolverID)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return solverId, nil
}

func (s *TaskStorage) SetIsActive(ctx context.Context, taskId int64) error {
	const op = "TaskStorage.SetIsActive"

	query := `UPDATE tasks SET is_active = NOT is_active WHERE id = $1;`

	_, err := s.db.ExecContext(ctx, query, taskId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TaskStorage) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "TaskStorage.Delete"

	rowsAffected, err = storage.Delete(ctx, s.db, "tasks", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
