package models

type Role struct {
	ID             int64    `db:"id"`
	Name           string   `db:"name"`
	OrganizationID int64    `db:"organization_id"`
	Permissions    []string `db:"permissions"`
}

type AddRolePayload struct {
	Name           string
	OrganizationID int64
	Permissions    []string
}

type UpdateRolePayload struct {
	ID             int64
	Name           string
	OrganizationID int64
	Permissions    []string
}
