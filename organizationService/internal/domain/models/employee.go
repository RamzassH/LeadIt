package models

type Employee struct {
	ID             int64 `db:"id"`
	UserID         int64 `db:"user_id"`
	OrganizationID int64 `db:"organization_id"`
}

type AddEmployee struct {
	UserID         int64
	OrganizationID int64
}

type UpdateEmployee struct {
	ID             int64
	UserID         int64
	OrganizationID int64
}
