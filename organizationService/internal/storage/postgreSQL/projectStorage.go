package postgreSQL

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
)

type ProjectStorage struct {
	db *sql.DB
}

func NewProjectStorage(db *sql.DB) (*ProjectStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &ProjectStorage{db: db}, nil
}

func (*ProjectStorage) SaveProject(ctx context.Context, payload models.AddProjectPayload) (int64, error) {
	panic("implement me")
}

func (*ProjectStorage) GetOrganization(ctx context.Context, id int64) (*models.Organization, error) {
	panic("implement me")
}

func (*ProjectStorage) GetAllOrganizations(ctx context.Context) ([]models.Organization, error) {
	panic("implement me")
}

func (*ProjectStorage) UpdateOrganization(ctx context.Context, payload models.UpdateProjectPayload) (int64, error) {
	panic("implement me")
}

func (*ProjectStorage) DeleteOrganization(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
