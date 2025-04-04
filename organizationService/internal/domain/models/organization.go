package models

type OrganizationDTO struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	OrganizerID int64  `db:"organizer_id"`
	Description string `db:"description"`
	Image       string `db:"image"`
}

type GetOrganizationDTO struct {
	OrganizationID int64
}

type GetOrganizationsDTO struct {
	OrganizerID int64
}

type CreateOrganizationDTO struct {
	Name              string
	Description       string
	OrganizationImage string
}

type UpdateOrganizationDTO struct {
	ID          int64
	Name        string
	Description string
	Image       string
}
