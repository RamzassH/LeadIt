package models

import "time"

type CommentDTO struct {
	ID     int64
	TaskID int64
	Date   time.Time
	UserID int64
	Body   string
}

type CreateCommentDTO struct {
	TaskID int64
	UserID int64
	Body   string
}

type UpdateCommentDTO struct {
	ID     int64
	UserID int64
	Body   string
}
