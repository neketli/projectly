package usecase

import (
	"context"
	"task-tracker-server/internal/domain/project/entity"
	"task-tracker-server/internal/domain/project/repository"
	"task-tracker-server/pkg/logger"
)

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

func New(r repository.ProjectRepository, l *logger.Logger) ProjectUseCase {
	return &projectUseCase{
		repo:   r,
		logger: l,
	}
}
