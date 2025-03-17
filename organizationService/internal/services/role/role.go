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
	GetRoleById(ctx context.Context, id int64) (role *models.Role, err error)
	GetAllRoles(ctx context.Context, organizationId int64) (roles []models.Role, err error)
	UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (role *models.Role, err error)
	DeleteRole(ctx context.Context, id int64) (rowsAffected int64, err error)
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
	const op = "role.AddRole"

	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("adding role")

	roleId, err := r.roleSaver.SaveRole(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Msg("failed to save role")
		return 0, err
	}

	return roleId, nil
}

func (r *Role) GetRole(ctx context.Context, id int64) (*models.Role, error) {
	const op = "role.GetRole"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting role")

	role, err := r.roleProvider.GetRoleById(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get role")
		return nil, err
	}

	return role, nil
}
func (r *Role) GetAllRoles(ctx context.Context, organizationId int64) ([]models.Role, error) {
	const op = "role.GetAllRoles"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all roles")
	roles, err := r.roleProvider.GetAllRoles(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all roles")
		return nil, err
	}

	return roles, nil
}
func (r *Role) UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (*models.Role, error) {
	const op = "role.UpdateRole"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("updating role")

	updatedRole, err := r.roleProvider.UpdateRole(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Msg("failed to update role")
		return nil, err
	}

	return updatedRole, nil
}
func (r *Role) DeleteRole(ctx context.Context, id int64) (int64, error) {
	const op = "role.DeleteRole"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("deleting role")

	deletedRole, err := r.roleProvider.DeleteRole(ctx, id)

	if err != nil {
		logger.Error().Err(err).Msg("failed to delete role")
		return 0, err
	}
	return deletedRole, nil
}
