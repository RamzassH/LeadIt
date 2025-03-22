package models

type Project struct {
	ID             int64
	Name           string
	Description    string
	OrganizationID int64
	Image          string
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
