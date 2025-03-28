package postgreSQL

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/RamzassH/LeadIt/organizationService/internal/domain/models"
	"github.com/RamzassH/LeadIt/organizationService/internal/storage"
	"github.com/jmoiron/sqlx"
	pq "github.com/lib/pq"
)

type ProjectStorage struct {
	db *sqlx.DB
}

func NewProjectStorage(db *sqlx.DB) (*ProjectStorage, error) {
	if db == nil {
		return nil, fmt.Errorf("datatabase connection is nil")
	}

	return &ProjectStorage{db: db}, nil
}

func (s *ProjectStorage) SaveProject(ctx context.Context, payload models.AddProjectPayload) (projectId int64, err error) {
	const op = "storage.saveProject"

	query := `
				INSERT INTO projects (name, description, organization_id, project_image)
				VALUES ($1, $2, $3, $4)
				RETURNING id`

	err = s.db.QueryRowContext(ctx, query, payload.Name, payload.Description, payload.OrganizationID, payload.Image).Scan(&projectId)

	if err != nil {
		var PgErr *pq.Error
		if errors.As(err, &PgErr) {
			if PgErr.Code.Name() == "unique_violation" {
				return 0, fmt.Errorf("%s %w", op, storage.ErrAlreadyExists)
			}
		}
		return 0, fmt.Errorf("%s %w", op, err)
	}

	return projectId, nil
}

func (s *ProjectStorage) GetProjectById(ctx context.Context, id int64) (project *models.Project, err error) {
	const op = "storage.getOrganization"

	err = storage.GetById(ctx, s.db, "projects", id, &project)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return project, nil
}

func (s *ProjectStorage) GetAllProjects(ctx context.Context, organizationId int64) (projects []models.Project, err error) {
	const op = "storage.getAllProjects"

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM projects WHERE organization_id = $1`, organizationId)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.OrganizationID); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func (s *ProjectStorage) UpdateProject(ctx context.Context, payload models.UpdateProjectPayload) (project *models.Project, err error) {
	const op = "storage.updateProject"

	query := `
			UPDATE projects
			SET 
			    name = COALESCE($1, name),
			    description = COALESCE($2, description),
			    organization_id = COALESCE($3, organization_id),
			    project_image = COALESCE($4, project_image)
			WHERE id = $5
			RETURNING id, name, description, organization_id, project_image`

	row := s.db.QueryRowContext(ctx, query, payload.Name, payload.Description, payload.OrganizationID, payload.Image)

	err = row.Scan(&project.ID, &project.Name, &project.Description, &project.OrganizationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return project, nil
}

func (s *ProjectStorage) DeleteProject(ctx context.Context, id int64) (rowsAffected int64, err error) {
	const op = "storage.deleteProject"

	rowsAffected, err = storage.Delete(ctx, s.db, "projects", id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}
