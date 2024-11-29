package project

import (
	"task-tracker-server/internal/domain/project/delivery"
	"task-tracker-server/internal/domain/project/repository"
	"task-tracker-server/internal/domain/project/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
	Router   *echo.Group
}

func New(dependency Dependency) usecase.ProjectUseCase {
	repo := repository.New(dependency.Postgres)

	projectUseCase := usecase.New(repo, dependency.Logger)
	delivery.New(dependency.Router, projectUseCase)

	return projectUseCase
}
