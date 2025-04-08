package role

import (
	"context"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Role struct {
	logger       zerolog.Logger
	roleSaver    Saver
	roleProvider Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateRoleDTO) (int64, error)
}
type Provider interface {
	GetById(ctx context.Context, id int64) (role *models.RoleDTO, err error)
	GetManyByOrganizationId(ctx context.Context, organizationId int64) (roles []models.RoleDTO, err error)
	Update(ctx context.Context, payload models.UpdateRoleDTO) (role *models.RoleDTO, err error)
	Delete(ctx context.Context, id int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	roleSaver Saver,
	roleProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer) *Role {
	return &Role{
		logger:       logger,
		roleSaver:    roleSaver,
		roleProvider: roleProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (r *Role) CreateRole(ctx context.Context, payload models.CreateRoleDTO) (int64, error) {
	const op = "role.CreateRole"

	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("adding role")

	roleId, err := r.roleSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Msg("failed to save role")
		return 0, err
	}

	return roleId, nil
}

func (r *Role) GetRole(ctx context.Context, id int64) (*models.RoleDTO, error) {
	const op = "role.GetRole"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting role")

	role, err := r.roleProvider.GetById(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get role")
		return nil, err
	}

	return role, nil
}
func (r *Role) GetAllRoles(ctx context.Context, organizationId int64) ([]models.RoleDTO, error) {
	const op = "role.GetAllRoles"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("getting all roles")
	roles, err := r.roleProvider.GetManyByOrganizationId(ctx, organizationId)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all roles")
		return nil, err
	}

	return roles, nil
}
func (r *Role) UpdateRole(ctx context.Context, payload models.UpdateRoleDTO) (*models.RoleDTO, error) {
	const op = "role.UpdateRole"
	logger := r.logger.With().Str("operation", op).Logger()
	logger.Info().Msg("updating role")

	updatedRole, err := r.roleProvider.Update(ctx, payload)
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

	deletedRole, err := r.roleProvider.Delete(ctx, id)

	if err != nil {
		logger.Error().Err(err).Msg("failed to delete role")
		return 0, err
	}
	return deletedRole, nil
}
