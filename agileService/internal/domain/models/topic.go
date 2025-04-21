package models

type TopicDTO struct {
	ID          int64  `json:"id"`
	ProjectID   int64  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTopicDTO struct {
	ProjectID   int64  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTopicDTO struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
