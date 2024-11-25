package repository

import (
	"context"
	"task-tracker-server/internal/domain/project/entity"
	"task-tracker-server/pkg/postgres"
	"time"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *entity.Project) error
	UpdateProject(ctx context.Context, project *entity.Project) error
	DeleteProject(ctx context.Context, projectID int) error
	GetProjectByCode(ctx context.Context, teamID int, code string) (entity.Project, error)
	GetProject(ctx context.Context, projectID int) (entity.Project, error)
	GetProjectList(ctx context.Context, teamID int) ([]entity.Project, error)
	GetUserProjects(ctx context.Context, userID int) ([]entity.Project, error)
}

const (
	_defaultConnTimeout = 5 * time.Second
)

type projectRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) ProjectRepository {
	return projectRepo{pg}
}
