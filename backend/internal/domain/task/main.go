package task

import (
	"projectly-server/internal/domain/task/delivery"
	"projectly-server/internal/domain/task/repository"
	taskUseCase "projectly-server/internal/domain/task/usecase"
	teamUseCase "projectly-server/internal/domain/team/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/minio"
	"projectly-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

// Dependency contains dependencies for task domain initialization.
type Dependency struct {
	Logger      *logger.Logger
	Postgres    *postgres.Postgres
	S3          *minio.Minio
	Router      *echo.Group
	TeamUseCase teamUseCase.TeamUseCase
}

// New initializes the task domain with its dependencies and returns a TaskUseCase.
func New(dependency Dependency) taskUseCase.TaskUseCase {
	repo := repository.New(dependency.Postgres, dependency.S3)

	usecase := taskUseCase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, usecase, dependency.TeamUseCase)

	return usecase
}
