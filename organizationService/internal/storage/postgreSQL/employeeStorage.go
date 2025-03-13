package postgreSQL

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
)

type EmployeeStorage struct {
	db *sql.DB
}

func NewEmployeeStorage(db *sql.DB) (*EmployeeStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &EmployeeStorage{db: db}, nil
}

func (*EmployeeStorage) SaveEmployee(ctx context.Context, payload models.AddEmployee) (int64, error) {
	panic("implement me")
}
func (*EmployeeStorage) GetEmployeeById(ctx context.Context, id int64) (*models.Employee, error) {
	panic("implement me")
}

func (*EmployeeStorage) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	panic("implement me")
}

func (*EmployeeStorage) UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (int64, error) {
	panic("implement me")
}

func (*EmployeeStorage) DeleteEmployee(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
