package models

type Role struct {
	ID             int64
	Name           string
	OrganizationID int64
	Permissions    []string
}

type AddRolePayload struct {
	Name        string
	Description string
	Image       string
}

type UpdateRolePayload struct {
	ID             int64
	Name           string
	OrganizationID int64
	Permissions    []string
}
