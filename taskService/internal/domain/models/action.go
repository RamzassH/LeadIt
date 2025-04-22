package models

import "time"

type ActionDTO struct {
	ID          int64     `json:"id"`
	TaskID      int64     `json:"task_id"`
	UserID      int64     `json:"user_id"`
	ActionType  string    `json:"action_type"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type CreateActionDTO struct {
	TaskID      int64  `json:"task_id"`
	UserID      int64  `json:"user_id"`
	ActionType  string `json:"action_type"`
	Description string `json:"description"`
}
