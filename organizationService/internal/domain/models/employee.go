package models

type EmployeeDTO struct {
	ID             int64 `db:"id"`
	UserID         int64 `db:"user_id"`
	OrganizationID int64 `db:"organization_id"`
}

type CreateEmployeeDTO struct {
	UserID         int64
	OrganizationID int64
}

type UpdateEmployeeRoleDTO struct {
	ID     int64
	RoleID int64
}
