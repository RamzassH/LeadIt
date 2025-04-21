package models

type TaskTopicDTO struct {
	TaskID  int64 `json:"task_id"`
	TopicID int64 `json:"topic_id"`
}
