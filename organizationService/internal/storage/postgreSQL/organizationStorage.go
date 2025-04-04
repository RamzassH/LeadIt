package postgreSQL

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/grpc/interceptors"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type OrganizationStorage struct {
	db *sqlx.DB
}

func NewOrganizationStorage(db *sqlx.DB) (*OrganizationStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &OrganizationStorage{db: db}, nil
}

func (s *OrganizationStorage) Save(
	ctx context.Context,
	payload models.CreateOrganizationDTO) (organizationId int64, err error) {
	const op = "OrganizationStorage.Save"

	query := `
	INSERT INTO organizations (name, description, organizer_id, organization_image)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	organizerID, ok := ctx.Value(interceptors.CtxUserID).(int64)
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

func (s *OrganizationStorage) GetById(ctx context.Context, id int64) (*models.OrganizationDTO, error) {
	const op = "OrganizationStorage.GetById"

	var org models.OrganizationDTO
	err := storage.GetById(ctx, s.db, "organizations", id, &org)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &org, nil
}

func (s *OrganizationStorage) GetByName(ctx context.Context, name string) (*models.OrganizationDTO, error) {
	const op = "OrganizationStorage.GetByName"

	row := s.db.QueryRowContext(ctx,
		`SELECT id, name, description, organizer_id, organization_image 
        FROM organizations WHERE name = $1`,
		name,
	)

	var org models.OrganizationDTO
	err := row.Scan(
		&org.ID,
		&org.Name,
		&org.Description,
		&org.OrganizerID,
		&org.Image,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &org, nil
}

func (s *OrganizationStorage) GetManyByOrganizerId(ctx context.Context, organizerId int64) ([]models.OrganizationDTO, error) {
	const op = "OrganizationStorage.GetManyByOrganizerId"

	rows, err := s.db.QueryContext(ctx,
		`SELECT id, name, organizer_id, description, organization_image 
    			FROM organizations 
   				WHERE organizer_id=$1;`,
		organizerId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var organizations []models.OrganizationDTO

	for rows.Next() {
		var org models.OrganizationDTO
		if err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.OrganizerID,
			&org.Description,
			&org.Image,
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

func (s *OrganizationStorage) Update(ctx context.Context, payload models.UpdateOrganizationDTO) (*models.OrganizationDTO, error) {
	const op = "OrganizationStorage.Update"

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
		payload.Image,
		payload.ID,
	)

	var organization models.OrganizationDTO
	err := row.Scan(
		&organization.ID,
		&organization.Name,
		&organization.Description,
		&organization.OrganizerID,
		&organization.Image,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &organization, nil
}

func (s *OrganizationStorage) Delete(ctx context.Context, id int64) (int64, error) {
	const op = "OrganizationStorage.Delete"

	rowsAffected, err := storage.Delete(ctx, s.db, "organizations", id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
