package models

type Roles struct {
	ID             int64
	Name           string
	OrganizationID int64
	permissions    []string
}
