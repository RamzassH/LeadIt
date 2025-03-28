package models

type Project struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	OrganizationID int64  `db:"organization_id"`
	Image          string `json:"image"`
}

type AddProjectPayload struct {
	Name           string
	Description    string
	OrganizationID int64
	Image          string
}

type UpdateProjectPayload struct {
	ID             int64
	Name           string
	Description    string
	OrganizationID int64
	Image          string
}
