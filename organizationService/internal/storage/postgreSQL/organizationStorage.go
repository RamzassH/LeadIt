package postgreSQL

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage"
	"github.com/lib/pq"
)

type OrganizationStorage struct {
	db *sql.DB
}

func NewOrganizationStorage(db *sql.DB) (*OrganizationStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &OrganizationStorage{db: db}, nil
}

func (s *OrganizationStorage) SaveOrganization(
	ctx context.Context,
	payload models.AddOrganizationPayload) (organizationId int64, err error) {
	const op = "storage.saveOrganization"

	query := `
	INSERT INTO organizations (name, description, organizer_id, organization_image)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	organizerID, ok := ctx.Value("user_id").(int64)
	if !ok {
		return 0, fmt.Errorf("%s: %w", op, errors.New("missing organizer id in context"))
	}
	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.Description, organizerID, payload.OrganizationImage).Scan(&organizationId)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return organizationId, nil
}

func (s *OrganizationStorage) GetOrganizationById(ctx context.Context, id int64) (organization *models.Organization, err error) {
	const op = "storage.getOrganizationById"

	err = storage.GetById(ctx, s.db, "organizations", id, &organization)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return organization, nil
}
func (s *OrganizationStorage) GetAllOrganizations(ctx context.Context, organizerId int64) ([]models.Organization, error) {
	const op = "storage.getAllOrganizations"

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM organizations WHERE organizer_id=$1;`, organizerId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var organizations []models.Organization

	for rows.Next() {
		var org models.Organization
		if err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.Description,
			&org.OrganizerID,
			&org.OrganizationImage,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		organizations = append(organizations, org)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return organizations, nil
}

func (s *OrganizationStorage) UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (*models.Organization, error) {
	const op = "storage.updateOrganization"

	query := `
        UPDATE organizations
        SET 
            name = COALESCE($1, name),
            description = COALESCE($2, description),
            organization_image = COALESCE($3, organization_image)
        WHERE id = $4
        RETURNING id, name, description, organizer_id, organization_image`

	row := s.db.QueryRowContext(ctx, query,
		payload.Name,
		payload.Description,
		payload.OrganizationImage,
		payload.ID,
	)

	var organization models.Organization
	err := row.Scan(
		&organization.ID,
		&organization.Name,
		&organization.Description,
		&organization.OrganizerID,
		&organization.OrganizationImage,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &organization, nil
}

func (s *OrganizationStorage) DeleteOrganization(ctx context.Context, id int64) (int64, error) {
	const op = "storage.DeleteOrganization"

	rowsAffected, err := storage.Delete(ctx, s.db, "organizations", id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}

func (s *OrganizationStorage) GetOrganizationByName(ctx context.Context, name string) (*models.Organization, error) {
	const op = "storage.GetOrganizationByName"

	row := s.db.QueryRowContext(ctx,
		`SELECT id, name, description, organizer_id, organization_image 
        FROM organizations WHERE name = $1`,
		name,
	)

	var org models.Organization
	err := row.Scan(
		&org.ID,
		&org.Name,
		&org.Description,
		&org.OrganizerID,
		&org.OrganizationImage,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &org, nil
}
