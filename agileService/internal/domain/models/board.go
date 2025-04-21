package models

type BoardDTO struct {
	ID          int64  `json:"id"`
	ProjectID   int64  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type CreateBoardDTO struct {
	ProjectID   int64  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type UpdateBoardDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
