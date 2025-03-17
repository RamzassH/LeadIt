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
	GetEmployeeById(ctx context.Context, id int64) (employee *models.Employee, err error)
	GetAllEmployees(ctx context.Context, organizationId int64) (employees []models.Employee, err error)
	UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (employee *models.Employee, err error)
	DeleteEmployee(ctx context.Context, id int64) (rowsAffected int64, err error)
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
	const op = "employee.AddEmployee"
	logger := e.logger.With().Str("operation", "AddEmployee").Logger()

	logger.Info().Str("operation", op).Msg("adding employee")
	employee, err := e.employeeSaver.SaveEmployee(ctx, payload)

	if err != nil {
		logger.Error().Err(err).Str("operation", "SaveEmployee").Msg(err.Error())
		return 0, err
	}

	return employee, nil
}

func (e *Employee) GetEmployee(ctx context.Context, id int64) (*models.Employee, error) {
	const op = "employee.GetEmployee"
	logger := e.logger.With().Int64("employeeId", id).Logger()

	logger.Info().Str("operation", op).Msg("getting employee")
	employee, err := e.employeeProvider.GetEmployeeById(ctx, id)

	if err != nil {
		logger.Error().Err(err).Str("operation", "GetEmployee").Msg(err.Error())
		return nil, err
	}

	return employee, nil
}

func (e *Employee) GetAllEmployees(ctx context.Context, organizationId int64) ([]models.Employee, error) {
	const op = "employee.GetAllEmployees"
	logger := e.logger.With().Str("operation", "GetAllEmployees").Logger()
	logger.Info().Str("operation", op).Msg("getting all employees")

	employee, err := e.employeeProvider.GetAllEmployees(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Str("operation", "GetAllEmployees").Msg(err.Error())
		return nil, err
	}
	return employee, nil
}

func (e *Employee) UpdateEmployee(ctx context.Context, payload models.UpdateEmployee) (*models.Employee, error) {
	const op = "employee.UpdateEployee"
	logger := e.logger.With().Str("operation", "UpdateEmployee").Logger()
	logger.Info().Msg("updating employee")
	employee, err := e.employeeProvider.UpdateEmployee(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", "UpdateEmployee").Msg(err.Error())
		return nil, err
	}

	return employee, nil
}

func (e *Employee) DeleteEmployee(ctx context.Context, id int64) (int64, error) {
	const op = "employee.DeleteEmployee"
	logger := e.logger.With().Int64("employeeId", id).Logger()
	logger.Info().Msg("deleting employee")

	deletedId, err := e.employeeProvider.DeleteEmployee(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", "DeleteEmployee").Msg(err.Error())
		return 0, err
	}

	return deletedId, nil
}
