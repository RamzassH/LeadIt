package models

type StatusDTO struct {
	ID             int64
	Name           string
	Order          int32
	ProgressBarID  int64
	NotifyRolesIDS []int64
}

type CreateStatusDTO struct {
	Name           string  `json:"name"`
	Order          int32   `json:"order"`
	ProgressBarID  int64   `json:"progressBarID"`
	NotifyRolesIDS []int64 `json:"notifyRolesIDS"`
}

type UpdateStatusDTO struct {
	ID             int64
	Name           string
	Order          int32
	NotifyRolesIDS []int64
}
