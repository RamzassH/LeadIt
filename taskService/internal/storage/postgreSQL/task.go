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

func (s *TaskStorage) GetManyByProjectId(ctx context.Context, projectId int64) (tasList []*models.TaskDTO, err error) {
	const op = "Task.GetManyByProjectId"

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM tasks WHERE project_id = $1`, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.TaskDTO
		if err := rows.Scan(item.ID, item.Name, item.Description, item.ProjectID, item.SetterID, item.SolverID,
			item.AllocatedTime, item.LastStartTime, item.SprintID, item.WastedTime, item.ProgressBarID, item.CurrentStatusID); err != nil {
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
                	current_status_id = COALESCE($9, current_status_id),
                 WHERE id = $1 RETURNING *;`
	var updated models.TaskDTO
	err := s.db.GetContext(ctx, &updated, query, payload.Name, payload.Description,
		payload.SolverID, payload.AllocatedTime,
		payload.SprintID, payload.WastedTime,
		payload.ProgressBarID, payload.CurrentStatusID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &updated, nil
}

func (s *TaskStorage) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "TaskStorage.Delete"

	rowsAffected, err = storage.Delete(ctx, s.db, "tasks", id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return rowsAffected, nil
}
