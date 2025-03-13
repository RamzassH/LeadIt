package organization

import (
	"context"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
	"kafka"
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

	GetAllOrganizations(ctx context.Context) ([]models.Organization, error)

	UpdateOrganizationBy(ctx context.Context, payload models.UpdateOrganizationPayload) (int64, error)

	DeleteOrganizationBy(ctx context.Context, id int64) (int64, error)
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
	panic("implement me")
}
func (org *Organization) GetOrganization(ctx context.Context, payload models.GetOrganizationPayload) (*models.Organization, error) {
	panic("implement me")
}
func (org *Organization) GetAllOrganizations(ctx context.Context, payload models.GetOrganizationsPayload) ([]models.Organization, error) {
	panic("implement me")
}

func (org *Organization) UpdateOrganization(ctx context.Context, payload models.UpdateOrganizationPayload) (models.Organization, error) {
	panic("implement me")
}

func (org *Organization) DeleteOrganization(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
