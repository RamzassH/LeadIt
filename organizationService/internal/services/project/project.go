package project

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
	"time"
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

	key := fmt.Sprintf("project:%d", id)

	cached, err := p.redisStorage.Get(ctx, key)
	if err != nil {
		var data models.ProjectDTO
		if err := json.Unmarshal([]byte(cached), &data); err == nil {
			logger.Debug().Int64("projectId", id).Msg("project loaded from cache")
			return &data, nil
		}
	}

	project, err := p.projectProvider.GetById(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get project")
		return nil, err
	}

	if data, err := json.Marshal(&project); err == nil {
		err := p.redisStorage.Set(ctx, key, data, 120*time.Minute)
		if err != nil {
			logger.Warn().Err(err).Str("operation", op).Msg("failed to save project")
		}
	} else {
		logger.Warn().Err(err).Str("operation", op).Msg("failed to save project in cache")
	}

	return project, nil
}
func (p *Project) GetAllProjects(ctx context.Context, organizationId int64) ([]models.ProjectDTO, error) {
	const op = "project.GetAllProjects"

	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all projects")

	key := fmt.Sprintf("organization:%d:projects", organizationId)

	cached, err := p.redisStorage.Get(ctx, key)
	if err == nil && cached != "" {
		var projects []models.ProjectDTO
		if err := json.Unmarshal([]byte(cached), &projects); err == nil {
			return projects, nil
		}
		logger.Warn().Err(err).Msg("failed to unmarshal cached task list")
	}

	projects, err := p.projectProvider.GetManyByOrganizationId(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get all projects")
		return nil, err
	}

	if data, err := json.Marshal(projects); err == nil {
		err = p.redisStorage.Set(ctx, key, data, 120*time.Minute)
		if err != nil {
			logger.Error().Err(err).Str("operation", op).Msg("failed to save projects")
		}
	} else {
		logger.Warn().Err(err).Msg("failed to marshal projects")
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

	p.RefreshProjectCache(ctx, updatedProject.ID)

	return updatedProject, nil
}

func (p *Project) DeleteProject(ctx context.Context, id int64) (int64, error) {
	const op = "project.DeleteProject"

	logger := p.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("deleting project")

	rowsAffected, err := p.projectProvider.Delete(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to delete project")
		return 0, err
	}

	key := fmt.Sprintf("project:%d", id)

	_ = p.redisStorage.Del(ctx, key)
	return rowsAffected, err
}
