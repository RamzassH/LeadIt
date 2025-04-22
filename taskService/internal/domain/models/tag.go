package models

type TagDTO struct {
	ID        int64
	Name      string
	TagColor  string
	ProjectID int64
}

type CreateTagDTO struct {
	Name      string `json:"name"`
	TagColor  string `json:"tag_color"`
	ProjectID int64  `json:"project_id"`
}

type UpdateTagDTO struct {
	ID       int64
	Name     string `json:"name"`
	TagColor string `json:"tag_color"`
}
