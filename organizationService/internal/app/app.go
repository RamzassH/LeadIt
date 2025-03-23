package app

import (
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/app/grpc"
	"github.com/RamzassH/LeadIt/organizationService/internal/config"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/employee"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/organization"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/project"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/role"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage/postgreSQL"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

type Services struct {
	organization.Organization
	role.Role
	project.Project
	employee.Employee
}

func New(logger zerolog.Logger,
	config *config.Config,
	validate *validator.Validate,
	db *sql.DB,
	redisClient *redis.Client,
	kafka *kafka.Producer) (*App, error) {

	organizationStorage, err := postgreSQL.NewOrganizationStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize organization storage: %w", err)
	}
	projectStorage, err := postgreSQL.NewProjectStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize project storage: %w", err)
	}
	roleStorage, err := postgreSQL.NewRoleStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize role storage: %w", err)
	}
	employeeStorage, err := postgreSQL.NewEmployeeStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize employee storage: %w", err)
	}

	rStorage, redisErr := redisStorage.New(redisClient, logger)

	if redisErr != nil {
		return nil, fmt.Errorf("failed to initialize redis storage: %w", redisErr)
	}

	organizationService := organization.New(
		logger,
		organizationStorage,
		organizationStorage,
		rStorage,
		kafka,
	)
	roleService := role.New(
		logger,
		roleStorage,
		roleStorage,
		rStorage,
		kafka)
	projectService := project.New(
		logger,
		projectStorage,
		projectStorage,
		rStorage,
		kafka)
	employeeService := employee.New(
		logger,
		employeeStorage,
		employeeStorage,
		rStorage,
		kafka)

	services := &Services{
		Organization: *organizationService,
		Role:         *roleService,
		Project:      *projectService,
		Employee:     *employeeService,
	}

	grpcApp := grpcapp.New(logger, config, validate, services)

	return &App{
		GRPCServer: grpcApp,
		validator:  validate,
	}, nil
}
