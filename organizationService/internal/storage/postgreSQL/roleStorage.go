package postgreSQL

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
)

type RoleStorage struct {
	db *sql.DB
}

func NewRoleStorage(db *sql.DB) (*RoleStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &RoleStorage{db: db}, nil
}

func (s *RoleStorage) SaveRole(ctx context.Context, payload models.AddRolePayload) (int64, error) {
	panic("implement me")
}

func (s *RoleStorage) GetRoleById(ctx context.Context, id int64) (*models.Role, error) {
	panic("implement me")
}
func (s *RoleStorage) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	panic("implement me")
}
func (s *RoleStorage) UpdateRole(ctx context.Context, payload models.UpdateRolePayload) (int64, error) {
	panic("implement me")
}
func (s *RoleStorage) DeleteRole(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
