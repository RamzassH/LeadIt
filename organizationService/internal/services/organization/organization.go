package organization

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage"
	"github.com/rs/zerolog"
	"time"
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
	const op = "organization.CreateOrganization"

	logger := org.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("Creating organization")
	organization, err := org.organizationProvider.GetByName(ctx, payload.Name)
	if err != nil {
		if !errors.Is(err, storage.ErrNotFound) {
			logger.Error().Err(err).Msg(err.Error())
			return 0, err
		}
	}
	if organization != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	organizationId, err := org.organizationSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Msg("failed to save organization")
		return 0, err
	}

	_ = org.redisStorage.Del(ctx, org.OrgListKey(payload.OrganizerID))

	return organizationId, nil

}

func (org *Organization) GetOrganization(ctx context.Context, payload models.GetOrganizationDTO) (*models.OrganizationDTO, error) {
	const op = "organization.GetOrganization"
	logger := org.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("getting organization")

	key := org.OrgKey(payload.OrganizationID)

	cachedOrganization, err := org.redisStorage.Get(ctx, key)
	if err == nil && cachedOrganization != "" {
		var data models.OrganizationDTO
		if err := json.Unmarshal([]byte(cachedOrganization), &data); err == nil {
			logger.Debug().Int64("organizationId", payload.OrganizationID).Msg("organization loaded from cache")
			return &data, nil
		}
	}

	organization, err := org.organizationProvider.GetById(ctx, payload.OrganizationID)

	if data, err := json.Marshal(organization); err == nil {
		err = org.redisStorage.Set(ctx, key, data, 120*time.Minute)
		if err != nil {
			logger.Warn().Err(err).Msg("failed to save organization")
		}
	} else {
		logger.Warn().Err(err).Msg("failed to marshal organization")
	}

	if err != nil {
		logger.Error().Err(err).Msg("failed to get organization")
		return nil, err
	}

	return organization, nil
}

func (org *Organization) GetAllOrganizations(ctx context.Context, payload models.GetOrganizationsDTO) ([]models.OrganizationDTO, error) {
	const op = "organization.GetAllOrganizations"

	logger := org.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all organizations")

	key := org.OrgListKey(payload.OrganizerID)

	cached, err := org.redisStorage.Get(ctx, key)
	if err == nil && cached != "" {
		var organizations []models.OrganizationDTO
		if err := json.Unmarshal([]byte(cached), &organizations); err == nil {
			return organizations, nil
		}
		logger.Warn().Err(err).Msg("failed to unmarshal cached task list")
	}

	organizations, err := org.organizationProvider.GetManyByOrganizerId(ctx, payload.OrganizerID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all organizations")
		return nil, err
	}

	if data, err := json.Marshal(organizations); err == nil {
		err = org.redisStorage.Set(ctx, key, data, 120*time.Minute)
		if err != nil {
			logger.Error().Err(err).Msg("failed to save organizations")
		}
	} else {
		logger.Warn().Err(err).Msg("failed to marshal organizations")
	}

	return organizations, nil
}

func (org *Organization) UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationDTO) (*models.OrganizationDTO, error) {
	const op = "organization.UpdateOrganization"
	logger := org.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("updating organization")

	organization, err := org.organizationProvider.Update(ctx, payload)

	if err != nil {
		logger.Error().Err(err).Msg("failed to update organization")
		return nil, err
	}

	org.RefreshOrganizationCache(ctx, organization.ID)

	return organization, err
}

func (org *Organization) DeleteOrganization(ctx context.Context, organizationID, organizerID int64) (int64, error) {
	const op = "organization.DeleteOrganization"
	logger := org.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("deleting organization")

	organizationID, err := org.organizationProvider.Delete(ctx, organizationID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to delete organization")
		return 0, err
	}

	_ = org.redisStorage.Del(ctx, org.OrgKey(organizationID))
	_ = org.redisStorage.Del(ctx, org.OrgListKey(organizerID))

	return organizationID, nil
}
