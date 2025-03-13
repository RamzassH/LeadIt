package project

import (
	"context"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
	"kafka"
	"time"
)

type Project struct {
	logger          zerolog.Logger
	projectSaver    Saver
	projectProvider Provider
	redisStorage    Redis
	kafka           *kafka.Producer
}

type Saver interface {
	SaveProject(ctx context.Context, payload models.AddProjectPayload) (int64, error)
}
type Provider interface {
	GetOrganization(ctx context.Context, id int64) (*models.Organization, error)

	GetAllOrganizations(ctx context.Context) ([]models.Organization, error)

	UpdateOrganization(ctx context.Context, payload models.UpdateProjectPayload) (int64, error)

	DeleteOrganization(ctx context.Context, id int64) (int64, error)
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
	projectSaver Saver,
	projectProvider Provider,
	redisStorage Redis,
	kafka *kafka.Producer) *Project {
	return &Project{
		logger:          logger,
		projectSaver:    projectSaver,
		projectProvider: projectProvider,
		redisStorage:    redisStorage,
		kafka:           kafka,
	}
}
func (p *Project) AddProject(ctx context.Context, payload models.AddProjectPayload) (int64, error) {
	panic("implement me")
}

func (p *Project) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	panic("implement me")
}
func (p *Project) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	panic("implement me")
}
func (p *Project) UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (int64, error) {
	panic("implement me")
}
func (p *Project) DeleteProject(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
