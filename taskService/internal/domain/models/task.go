package models

import "time"

type TaskDTO struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ProjectID       int64     `json:"project_id"`
	SetterID        int64     `json:"setter_id"`
	SolverID        int64     `json:"solver_id"`
	ProgressBarID   int64     `json:"progress_bar_id"`
	CurrentStatusID int64     `json:"current_status_id"`
	AllocatedTime   int64     `json:"allocated_time"`
	WastedTime      int64     `json:"wasted_time"`
	LastStartTime   time.Time `json:"last_start_time"`
	SprintID        int64     `json:"sprint_id"`
}

type CreateTaskDTO struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ProjectID       int64  `json:"project_id"`
	SetterID        int64  `json:"setter_id"`
	SolverID        int64  `json:"solver_id"`
	ProgressBarID   int64  `json:"progress_bar_id"`
	CurrentStatusID int64  `json:"current_status_id"`
	AllocatedTime   int64  `json:"allocated_time"`
	SprintID        int64  `json:"sprint_id"`
}

type UpdateTaskDTO struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	SolverID        int64  `json:"solver_id"`
	AllocatedTime   int64  `json:"allocated_time"`
	SprintID        int64  `json:"sprint_id"`
	WastedTime      int64  `json:"wasted_time"`
	ProgressBarID   int64  `json:"progress_bar_id"`
	CurrentStatusID int64  `json:"current_status_id"`
}
