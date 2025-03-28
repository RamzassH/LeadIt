package models

type Organization struct {
	ID                int64  `db:"id"`
	Name              string `db:"name"`
	OrganizerID       int64  `db:"organizer_id"`
	Description       string `db:"description"`
	OrganizationImage string `db:"organization_image"`
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
