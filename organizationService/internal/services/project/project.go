package project

import (
	"context"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Project struct {
	logger          zerolog.Logger
	projectSaver    Saver
	projectProvider Provider
	redisStorage    redisStorage.RedisStore
	kafka           *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateProjectDTO) (int64, error)
}
type Provider interface {
	GetById(ctx context.Context, id int64) (project *models.ProjectDTO, err error)

	GetManyByOrganizationId(ctx context.Context, organizationId int64) (projects []models.ProjectDTO, err error)

	Update(ctx context.Context, payload models.UpdateProjectDTO) (*models.ProjectDTO, error)

	Delete(ctx context.Context, id int64) (int64, error)
}

func New(
	logger zerolog.Logger,
	projectSaver Saver,
	projectProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer) *Project {
	return &Project{
		logger:          logger,
		projectSaver:    projectSaver,
		projectProvider: projectProvider,
		redisStorage:    redisStorage,
		kafka:           kafka,
	}
}
func (p *Project) CreateProject(ctx context.Context, payload models.CreateProjectDTO) (int64, error) {
	const op = "project.AddProject"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("adding project")

	projectId, err := p.projectSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to save project")
		return 0, err
	}

	return projectId, nil
}

func (p *Project) GetProject(ctx context.Context, id int64) (*models.ProjectDTO, error) {
	const op = "project.GetProject"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting project")

	project, err := p.projectProvider.GetById(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get project")
		return nil, err
	}

	return project, nil
}
func (p *Project) GetAllProjects(ctx context.Context, organizationId int64) ([]models.ProjectDTO, error) {
	const op = "project.GetAllProjects"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all projects")
	projects, err := p.projectProvider.GetManyByOrganizationId(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get all projects")
		return nil, err
	}

	return projects, nil
}
func (p *Project) UpdateProject(ctx context.Context, payload models.UpdateProjectDTO) (*models.ProjectDTO, error) {
	const op = "project.UpdateProject"
	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("updating project")

	updatedProject, err := p.projectProvider.Update(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to update project")
		return nil, err
	}

	return updatedProject, nil
}

func (p *Project) DeleteProject(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
