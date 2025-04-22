package employee

import (
	"context"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Employee struct {
	logger           zerolog.Logger
	employeeSaver    Saver
	employeeProvider Provider
	redisStorage     redisStorage.RedisStore
	kafka            *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateEmployeeDTO) (int64, error)
}
type Provider interface {
	GetById(ctx context.Context, id int64) (employee *models.EmployeeDTO, err error)
	GetManyByOrganizationId(ctx context.Context, organizationId int64) (employees []models.EmployeeDTO, err error)
	UpdateRole(ctx context.Context, payload models.UpdateEmployeeRoleDTO) (id int64, err error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	employeeSaver Saver,
	employeeProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer) *Employee {
	return &Employee{
		logger:           logger,
		employeeSaver:    employeeSaver,
		employeeProvider: employeeProvider,
		redisStorage:     redisStorage,
		kafka:            kafka,
	}
}

func (e *Employee) CreateEmployee(ctx context.Context, payload models.CreateEmployeeDTO) (int64, error) {
	const op = "employee.AddEmployee"
	logger := e.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("adding employee")
	employee, err := e.employeeSaver.Save(ctx, payload)

	if err != nil {
		logger.Error().Err(err).Msg(err.Error())
		return 0, err
	}

	return employee, nil
}

func (e *Employee) GetEmployee(ctx context.Context, id int64) (*models.EmployeeDTO, error) {
	const op = "employee.GetEmployee"
	logger := e.logger.With().Int64("employeeId", id).Logger()

	logger.Info().Str("operation", op).Msg("getting employee")
	employee, err := e.employeeProvider.GetById(ctx, id)

	if err != nil {
		logger.Error().Err(err).Str("operation", "GetEmployee").Msg(err.Error())
		return nil, err
	}

	return employee, nil
}

func (e *Employee) GetAllEmployees(ctx context.Context, organizationId int64) ([]models.EmployeeDTO, error) {
	const op = "employee.GetAllEmployees"
	logger := e.logger.With().Str("operation", "GetAllEmployees").Logger()
	logger.Info().Str("operation", op).Msg("getting all employees")

	employee, err := e.employeeProvider.GetManyByOrganizationId(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Str("operation", "GetAllEmployees").Msg(err.Error())
		return nil, err
	}
	return employee, nil
}

func (e *Employee) UpdateEmployeeRole(ctx context.Context, payload models.UpdateEmployeeRoleDTO) (int64, error) {
	const op = "employee.UpdateEmployee"
	logger := e.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("updating employee")
	employeeId, err := e.employeeProvider.UpdateRole(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", "UpdateEmployee").Msg(err.Error())
		return 0, err
	}

	return employeeId, nil
}

func (e *Employee) DeleteEmployee(ctx context.Context, id int64) (int64, error) {
	const op = "employee.DeleteEmployee"
	logger := e.logger.With().Str("operation", op).Int64("employeeId", id).Logger()
	logger.Info().Msg("deleting employee")

	deletedId, err := e.employeeProvider.Delete(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", "DeleteEmployee").Msg(err.Error())
		return 0, err
	}

	return deletedId, nil
}
