package models

type RoleDTO struct {
	ID             int64    `db:"id"`
	Name           string   `db:"name"`
	OrganizationID int64    `db:"organization_id"`
	Permissions    []string `db:"permissions"`
}

type CreateRoleDTO struct {
	Name           string
	OrganizationID int64
	Permissions    []string
}

type UpdateRoleDTO struct {
	ID          int64
	Name        string
	Permissions []string
}
