package project

import (
	"projectly-server/internal/domain/project/delivery"
	"projectly-server/internal/domain/project/repository"
	"projectly-server/internal/domain/project/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/postgres"

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
