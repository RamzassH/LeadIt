package models

type TaskBoardDTO struct {
	TaskID   int64 `json:"task_id"`
	BoardID  int64 `json:"board_id"`
	SprintID int64 `json:"sprint_id"`
	ColumnID int64 `json:"column_id"`
	Order    int32 `json:"order"`
}

type AttachTaskToBoardDTO struct {
	TaskID   int64 `json:"task_id"`
	BoardID  int64 `json:"board_id"`
	SprintID int64 `json:"sprint_id"`
	ColumnID int64 `json:"column_id"`
	Order    int32 `json:"order"`
}

type MoveTaskDTO struct {
	TaskID       int64 `json:"task_id"`
	FromColumnID int64 `json:"from_column_id"`
	ToColumnID   int64 `json:"to_column_id"`
	NewOrder     int32 `json:"new_order"`
	BoardID      int64 `json:"board_id"`
}
