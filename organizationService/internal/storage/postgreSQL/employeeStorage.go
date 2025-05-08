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

type EmployeeStorage struct {
	db *sqlx.DB
}

func NewEmployeeStorage(db *sqlx.DB) (*EmployeeStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &EmployeeStorage{db: db}, nil
}

func (s *EmployeeStorage) Save(ctx context.Context, payload models.CreateEmployeeDTO) (employeeID int64, err error) {
	const op = "EmployeeStorage.Save"

	query := `
	INSERT INTO employees (user_id, organization_id)
	VALUES ($1, $2)
	RETURNING id`

	err = s.db.QueryRowContext(ctx, query, payload.UserID, payload.OrganizationID).Scan(&employeeID)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return employeeID, nil
}

func (s *EmployeeStorage) GetById(ctx context.Context, id int64) (employee *models.EmployeeDTO, err error) {
	const op = "EmployeeStorage.GetById"

	err = storage.GetById(ctx, s.db, "employees", id, &employee)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return employee, nil
}

func (s *EmployeeStorage) GetManyByOrganizationId(ctx context.Context, organizationId int64) (employees []models.EmployeeDTO, err error) {
	const op = "EmployeeStorage.GetManyByOrganizationId"

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM employees WHERE organization_id = $1`, organizationId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.EmployeeDTO
		if err := rows.Scan(&employee.ID, &employee.UserID, &employee.OrganizationID); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return employees, nil
}

func (s *EmployeeStorage) UpdateRole(ctx context.Context, payload models.UpdateEmployeeRoleDTO) (id int64, err error) {
	const op = "EmployeeStorage.UpdateRole"

	query := `
		UPDATE employees
		SET 	
		    role = COALESCE($1, role)
		WHERE id = $2
		RETURNING id`

	row := s.db.QueryRowContext(ctx, query, payload.RoleID, payload.ID, payload.ID)

	err = row.Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, storage.ErrNotFound
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil

}

func (s *EmployeeStorage) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "EmployeeStorage.Delete"

	rowsAffected, err = storage.Delete(ctx, s.db, "employees", id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
