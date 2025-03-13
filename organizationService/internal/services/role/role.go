package role

import (
	"context"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
	"kafka"
	"time"
)

type Role struct {
	logger       zerolog.Logger
	roleSaver    Saver
	roleProvider Provider
	redisStorage Redis
	kafka        *kafka.Producer
}

type Saver interface {
	SaveRole(ctx context.Context, payload models.AddRolePayload) (int64, error)
}
type Provider interface {
	GetRoleById(ctx context.Context, id int64) (*models.Role, error)
	GetAllRoles(ctx context.Context) ([]models.Role, error)
	UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (int64, error)
	DeleteRole(ctx context.Context, id int64) (int64, error)
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
	roleSaver Saver,
	roleProvider Provider,
	redisStorage Redis,
	kafka *kafka.Producer) *Role {
	return &Role{
		logger:       logger,
		roleSaver:    roleSaver,
		roleProvider: roleProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (r *Role) AddRole(ctx context.Context, payload models.AddRolePayload) (int64, error) {
	panic("implement me")
}
func (r *Role) GetRole(ctx context.Context, id int64) (*models.Role, error) {
	panic("implement me")
}
func (r *Role) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	panic("implement me")
}
func (r *Role) UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (int64, error) {
	panic("implement me")
}
func (r *Role) DeleteRole(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
