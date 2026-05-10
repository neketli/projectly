package usecase

import (
	"context"
	"projectly-server/internal/domain/project/entity"
	"projectly-server/internal/domain/project/repository"
	"projectly-server/pkg/logger"
)

// ProjectUseCase defines the interface for project business logic.
type ProjectUseCase interface {
	CreateProject(ctx context.Context, project *entity.Project) error
	UpdateProject(ctx context.Context, project *entity.Project) error
	DeleteProject(ctx context.Context, projectID int) error
	GetProjectByCode(ctx context.Context, teamID int, code string) (entity.Project, error)
	GetProject(ctx context.Context, projectID int) (entity.Project, error)
	GetProjectList(ctx context.Context, teamID int) ([]entity.Project, error)
	GetUserProjects(ctx context.Context, userID int) ([]entity.Project, error)
}

type projectUseCase struct {
	repo   repository.ProjectRepository
	logger *logger.Logger
}

// New creates a new ProjectUseCase instance.
func New(r repository.ProjectRepository, l *logger.Logger) ProjectUseCase {
	return &projectUseCase{
		repo:   r,
		logger: l,
	}
}
