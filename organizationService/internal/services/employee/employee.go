package employee

import (
	"context"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
	"kafka"
	"time"
)

type Employee struct {
	logger           zerolog.Logger
	employeeSaver    Saver
	employeeProvider Provider
	redisStorage     Redis
	kafka            *kafka.Producer
}

type Saver interface {
	SaveEmployee(ctx context.Context, payload models.AddEmployee) (int64, error)
}
type Provider interface {
	GetEmployeeById(ctx context.Context, id int64) (*models.Employee, error)
	GetAllEmployees(ctx context.Context) ([]models.Employee, error)
	UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (int64, error)
	DeleteEmployee(ctx context.Context, id int64) (int64, error)
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}

func New(
	logger zerolog.Logger,
	employeeSaver Saver,
	employeeProvider Provider,
	redisStorage Redis,
	kafka *kafka.Producer) *Employee {
	return &Employee{
		logger:           logger,
		employeeSaver:    employeeSaver,
		employeeProvider: employeeProvider,
		redisStorage:     redisStorage,
		kafka:            kafka,
	}
}

func (e *Employee) AddEmployee(ctx context.Context, payload models.AddEmployee) (int64, error) {
	panic("implement me")
}

func (e *Employee) GetEmployee(ctx context.Context, id int64) (*models.Employee, error) {
	panic("implement me")
}

func (e *Employee) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	panic("implement me")
}

func (e *Employee) UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (int64, error) {
	panic("implement me")
}

func (e *Employee) DeleteEmployee(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
