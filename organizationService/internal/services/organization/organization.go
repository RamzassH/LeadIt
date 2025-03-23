package organization

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
	"time"
)

type Organization struct {
	logger               zerolog.Logger
	organizationSaver    Saver
	organizationProvider Provider
	redisStorage         Redis
	kafka                *kafka.Producer
}

type Saver interface {
	SaveOrganization(
		ctx context.Context,
		payload models.AddOrganizationPayload) (int64, error)
}
type Provider interface {
	GetOrganizationById(ctx context.Context, id int64) (*models.Organization, error)

	GetOrganizationByName(ctx context.Context, name string) (*models.Organization, error)

	GetAllOrganizations(ctx context.Context, organizerId int64) ([]models.Organization, error)

	UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (*models.Organization, error)

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
	organizationSaver Saver,
	organizationProvider Provider,
	redisStorage Redis,
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

func (org *Organization) AddOrganization(
	ctx context.Context,
	payload models.AddOrganizationPayload) (int64, error) {
	const op = "organization.AddOrganization"

	logger := org.logger.With().Str("operation", "AddOrganization").Logger()

	logger.Info().Str("operation", op).Msg("adding organization")
	organization, err := org.organizationProvider.GetOrganizationByName(ctx, payload.Name)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg(err.Error())
		return 0, err
	}
	if organization != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	organizationId, err := org.organizationSaver.SaveOrganization(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to save organization")
		return 0, err
	}

	return organizationId, nil

}

func (org *Organization) GetOrganization(ctx context.Context, payload models.GetOrganizationPayload) (*models.Organization, error) {
	const op = "organization.GetOrganization"
	logger := org.logger.With().Str("operation", "GetOrganization").Logger()

	logger.Info().Str("operation", op).Msg("getting organization")
	organization, err := org.organizationProvider.GetOrganizationById(ctx, payload.OrganizationID)

	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get organization")
		return nil, err
	}

	return organization, nil
}

func (org *Organization) GetAllOrganizations(ctx context.Context, payload models.GetOrganizationsPayload) ([]models.Organization, error) {
	const op = "organization.GetAllOrganizations"
	logger := org.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all organizations")
	organizations, err := org.organizationProvider.GetAllOrganizations(ctx, payload.OrganizerID)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to get all organizations")
		return nil, err
	}

	return organizations, nil
}

func (org *Organization) UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (*models.Organization, error) {
	const op = "organization.UpdateOrganization"
	logger := org.logger.With().Str("operation", "UpdateOrganization").Logger()
	logger.Info().Str("operation", op).Msg("updating organization")

	organization, err := org.organizationProvider.UpdateOrganization(ctx, payload)

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

	organizationID, err := org.organizationProvider.DeleteOrganization(ctx, id)
	if err != nil {
		logger.Error().Err(err).Str("operation", op).Msg("failed to delete organization")
		return 0, err
	}
	return organizationID, nil
}
