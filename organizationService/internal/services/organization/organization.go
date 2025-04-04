package organization

import (
	"context"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage"
	"github.com/rs/zerolog"
)

type Organization struct {
	logger               zerolog.Logger
	organizationSaver    Saver
	organizationProvider Provider
	redisStorage         redisStorage.RedisStore
	kafka                *kafka.Producer
}

type Saver interface {
	Save(
		ctx context.Context,
		payload models.CreateOrganizationDTO) (int64, error)
}
type Provider interface {
	GetById(ctx context.Context, id int64) (*models.OrganizationDTO, error)

	GetByName(ctx context.Context, name string) (*models.OrganizationDTO, error)

	GetManyByOrganizerId(ctx context.Context, organizerId int64) ([]models.OrganizationDTO, error)

	Update(ctx context.Context, payload models.UpdateOrganizationDTO) (*models.OrganizationDTO, error)

	Delete(ctx context.Context, id int64) (int64, error)
}

func New(
	logger zerolog.Logger,
	organizationSaver Saver,
	organizationProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *Organization {
	return &Organization{
		logger:               logger,
		organizationSaver:    organizationSaver,
		organizationProvider: organizationProvider,
		redisStorage:         redisStorage,
		kafka:                kafka,
	}
}

func (org *Organization) CreateOrganization(
	ctx context.Context,
	payload models.CreateOrganizationDTO) (int64, error) {
	const op = "organization.AddOrganization"

	logger := org.logger.With().Str("operation", "AddOrganization").Logger()

	logger.Info().Str("operation", op).Msg("adding organization")
	organization, err := org.organizationProvider.GetByName(ctx, payload.Name)
	if err != nil {
		if !errors.Is(err, storage.ErrNotFound) {
			logger.Error().Err(err).Str("operation", op).Msg(err.Error())
			return 0, err
		}
	}
	if organization != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	organizationId, err := org.organizationSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to save organization")
		return 0, err
	}

	return organizationId, nil

}

func (org *Organization) GetOrganization(ctx context.Context, payload models.GetOrganizationDTO) (*models.OrganizationDTO, error) {
	const op = "organization.GetOrganization"
	logger := org.logger.With().Str("operation", "GetOrganization").Logger()

	logger.Info().Str("operation", op).Msg("getting organization")
	organization, err := org.organizationProvider.GetById(ctx, payload.OrganizationID)

	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get organization")
		return nil, err
	}

	return organization, nil
}

func (org *Organization) GetAllOrganizations(ctx context.Context, payload models.GetOrganizationsDTO) ([]models.OrganizationDTO, error) {
	const op = "organization.GetAllOrganizations"
	logger := org.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all organizations")
	organizations, err := org.organizationProvider.GetManyByOrganizerId(ctx, payload.OrganizerID)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get all organizations")
		return nil, err
	}

	return organizations, nil
}

func (org *Organization) UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationDTO) (*models.OrganizationDTO, error) {
	const op = "organization.UpdateOrganization"
	logger := org.logger.With().Str("operation", "UpdateOrganization").Logger()
	logger.Info().Str("operation", op).Msg("updating organization")

	organization, err := org.organizationProvider.Update(ctx, payload)

	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to update organization")
		return nil, err
	}

	return organization, err
}

func (org *Organization) DeleteOrganization(ctx context.Context, id int64) (int64, error) {
	const op = "organization.DeleteOrganization"
	logger := org.logger.With().Str("operation", "DeleteOrganization").Logger()
	logger.Info().Str("operation", op).Msg("deleting organization")

	organizationID, err := org.organizationProvider.Delete(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to delete organization")
		return 0, err
	}
	return organizationID, nil
}
