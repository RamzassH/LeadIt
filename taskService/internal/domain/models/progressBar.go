package models

type ProgressBarDTO struct {
	ID        int64
	Name      string
	ProjectID int64
}

type CreateProgressBarDTO struct {
	Name      string `json:"name"`
	ProjectID int64  `json:"project_id"`
}

type UpdateProgressBarDTO struct {
	ID        int64
	Name      string `json:"name"`
	ProjectID int64  `json:"project_id"`
}
