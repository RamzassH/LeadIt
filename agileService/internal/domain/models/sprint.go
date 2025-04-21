package models

import "time"

type SprintDTO struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
}

type CreateSprintDTO struct {
	ProjectID   int64     `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type UpdateSprintDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
}
