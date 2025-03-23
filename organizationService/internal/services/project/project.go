package project

import (
	"context"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"

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
	GetProjectById(ctx context.Context, id int64) (project *models.Project, err error)

	GetAllProjects(ctx context.Context, organizationId int64) (projects []models.Project, err error)

	UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (*models.Project, error)

	DeleteProject(ctx context.Context, id int64) (int64, error)
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
	const op = "project.AddProject"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("adding project")

	projectId, err := p.projectSaver.SaveProject(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to save project")
		return 0, err
	}

	return projectId, nil
}

func (p *Project) GetProject(ctx context.Context, id int64) (*models.Project, error) {
	const op = "project.GetProject"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting project")

	project, err := p.projectProvider.GetProjectById(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get project")
		return nil, err
	}

	return project, nil
}
func (p *Project) GetAllProjects(ctx context.Context, organizationId int64) ([]models.Project, error) {
	const op = "project.GetAllProjects"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all projects")
	projects, err := p.projectProvider.GetAllProjects(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get all projects")
		return nil, err
	}

	return projects, nil
}
func (p *Project) UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (*models.Project, error) {
	const op = "project.UpdateProject"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("updating project")

	updatedProject, err := p.projectProvider.UpdateProject(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to update project")
		return nil, err
	}

	return updatedProject, nil
}
func (p *Project) DeleteProject(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
