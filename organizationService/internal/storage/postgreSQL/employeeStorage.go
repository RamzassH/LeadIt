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

func (s *EmployeeStorage) SaveEmployee(ctx context.Context, payload models.AddEmployee) (employeeID int64, err error) {
	const op = "storage.saveEmployee"

	query := `
	INSERT INTO employees (user_id, organization_id
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

func (s *EmployeeStorage) GetEmployeeById(ctx context.Context, id int64) (employee *models.Employee, err error) {
	const op = "storage.getEmployeeById"

	err = storage.GetById(ctx, s.db, "employees", id, &employee)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return employee, nil
}

func (s *EmployeeStorage) GetAllEmployees(ctx context.Context, organizationId int64) (employees []models.Employee, err error) {
	const op = "storage.getAllEmployees"

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM employees WHERE organization_id = $1`, organizationId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
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

func (s *EmployeeStorage) UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (employee *models.Employee, err error) {
	const op = "storage.updateEmployee"

	query := `
		UPDATE employees
		SET 	
		    user_id = COALESCE($1, user_id),
		    organization_id = COALESCE($2, organization_id)
		WHERE id = $3
		RETURNING id, user_id, organization_id`

	row := s.db.QueryRowContext(ctx, query, payload.UserID, payload.OrganizationID, payload.ID)

	err = row.Scan(&employee.ID, &employee.UserID, &employee.OrganizationID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return employee, nil

}

func (s *EmployeeStorage) DeleteEmployee(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "storage.deleteEmployee"

	rowsAffected, err = storage.Delete(ctx, s.db, "employees", id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
