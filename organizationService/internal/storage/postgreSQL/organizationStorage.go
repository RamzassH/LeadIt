package postgreSQL

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
)

type OrganizationStorage struct {
	db *sql.DB
}

func NewOrganizationStorage(db *sql.DB) (*OrganizationStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &OrganizationStorage{db: db}, nil
}

func (*OrganizationStorage) SaveOrganization(
	ctx context.Context,
	payload models.AddOrganizationPayload) (int64, error) {
	panic("implement me")
}

func (*OrganizationStorage) GetOrganizationById(ctx context.Context, id int64) (*models.Organization, error) {
	panic("implement me")
}

func (*OrganizationStorage) GetAllOrganizations(ctx context.Context) ([]models.Organization, error) {
	panic("implement me")
}

func (*OrganizationStorage) UpdateOrganizationBy(ctx context.Context, payload models.UpdateOrganizationPayload) (int64, error) {
	panic("implement me")
}

func (*OrganizationStorage) DeleteOrganizationBy(ctx context.Context, id int64) (int64, error) {
	panic("implement me")
}
