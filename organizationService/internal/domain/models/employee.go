package models

type Employee struct {
	ID             int64
	UserID         int64
	OrganizationID int64
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
