package app

import (
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/app/grpc"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/employee"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/facade"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/organization"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/project"
	"github.com/RamzassH/LeadIt/organizationService/internal/services/role"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage/postgreSQL"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"kafka"
	"redisStorage"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

func New(logger zerolog.Logger,
	grpcPort int,
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

	facadeService := facade.NewFacade(
		*organizationService,
		*roleService,
		*projectService,
		*employeeService)

	grpcApp := grpcapp.New(
		logger,
		grpcPort,
		validate,
		facadeService)

	return &App{
		GRPCServer: grpcApp,
		validator:  validate,
	}, nil
}
