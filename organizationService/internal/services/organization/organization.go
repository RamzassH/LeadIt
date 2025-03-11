package organization

import (
	"context"
	"github.com/rs/zerolog"
	"kafka"
	"time"
)

type Organization struct {
	logger               zerolog.Logger
	organizationSaver    OrganizationSaver
	organizationProvider OrganizationProvider
	roleSaver            RoleSaver
	roleProvider         RoleProvider
	employeeSaver        EmployeeSaver
	employeeProvider     EmployeeProvider
	projectSaver         ProjectSaver
	projectProvider      ProjectProvider
	redisStorage         Redis
	kafka                *kafka.Producer
}

type OrganizationSaver interface{}
type OrganizationProvider interface{}
type RoleSaver interface{}
type RoleProvider interface{}
type EmployeeSaver interface{}
type EmployeeProvider interface{}
type ProjectSaver interface{}
type ProjectProvider interface{}

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
	organizationSaver OrganizationSaver,
	organizationProvider OrganizationProvider,
	roleSaver RoleSaver,
	roleProvider RoleProvider,
	employeeSaver EmployeeSaver,
	employeeProvider EmployeeProvider,
	projectSaver ProjectSaver,
	projectProvider ProjectProvider,
	redisStorage Redis,
	kafka *kafka.Producer,
) *Organization {
	return &Organization{
		logger:               logger,
		organizationSaver:    organizationSaver,
		organizationProvider: organizationProvider,
		roleSaver:            roleSaver,
		roleProvider:         roleProvider,
		employeeSaver:        employeeSaver,
		employeeProvider:     employeeProvider,
		projectSaver:         projectSaver,
		projectProvider:      projectProvider,
		redisStorage:         redisStorage,
		kafka:                kafka,
	}
}
