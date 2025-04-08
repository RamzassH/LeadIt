package postgreSQL

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RoleStorage struct {
	db *sqlx.DB
}

func NewRoleStorage(db *sqlx.DB) (*RoleStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &RoleStorage{db: db}, nil
}

func (s *RoleStorage) Save(ctx context.Context, payload models.CreateRoleDTO) (roleId int64, err error) {
	const op = "RoleStorage.Save"

	query := `
			INSERT INTO roles (name, organization_id, permisstions)
			VALUES ($1, $2, $3)
			RETURNING id;`

	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.OrganizationID, payload.Permissions).Scan(&roleId)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return roleId, nil
}

func (s *RoleStorage) GetById(ctx context.Context, id int64) (role *models.RoleDTO, err error) {
	const op = "RoleStorage.GetById"

	err = storage.GetById(ctx, s.db, "roles", id, &role)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return role, nil
}

func (s *RoleStorage) GetManyByOrganizationId(ctx context.Context, organizationId int64) (roles []models.RoleDTO, err error) {
	const op = "RoleStorage.GetManyByOrganizationId"

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM roles WHERE id = $1`, organizationId)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var role models.RoleDTO
		if err := rows.Scan(&role.ID, &role.Name, &role.OrganizationID, &role.Permissions); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}
func (s *RoleStorage) Update(ctx context.Context, payload models.UpdateRoleDTO) (role *models.RoleDTO, err error) {
	const op = "RoleStorage.Update"

	query := `
			UPDATE roles
			SET 
			    name = COALESCE($1, name),
			    permisstions = COALESCE($2, permisstions)
			WHERE id = $3
			RETURNING id, name, organization_id, permisstions`

	row := s.db.QueryRowContext(ctx, query, payload.Name, payload.Permissions)

	err = row.Scan(&role.ID, &role.Name, &role.OrganizationID, &role.Permissions)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return role, nil
}
func (s *RoleStorage) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "RoleStorage.Delete"

	rowsAffected, err = storage.Delete(ctx, s.db, "roles", id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
