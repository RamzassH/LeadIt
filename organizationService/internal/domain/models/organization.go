package models

type Organization struct {
	ID                int64
	Name              string
	OrganizerID       int64
	Description       string
	OrganizationImage string
}

type GetOrganizationPayload struct {
	OrganizationID int64
}

type GetOrganizationsPayload struct {
	OrganizerID int64
}

type AddOrganizationPayload struct {
	Name              string
	Description       string
	OrganizationImage string
}

type UpdateOrganizationPayload struct {
	ID                int64
	Name              string
	OrganizerID       int64
	Description       string
	OrganizationImage string
}
