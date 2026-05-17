package project

import (
	"projectly-server/internal/domain/project/delivery"
	"projectly-server/internal/domain/project/repository"
	projUseCase "projectly-server/internal/domain/project/usecase"
	teamUseCase "projectly-server/internal/domain/team/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

// Dependency contains dependencies for project domain initialization.
type Dependency struct {
	Logger       *logger.Logger
	Postgres     *postgres.Postgres
	Router       *echo.Group
	TeamUseCase  teamUseCase.TeamUseCase
}

// New initializes the project domain with its dependencies and returns a ProjectUseCase.
func New(dependency Dependency) usecase.ProjectUseCase {
	repo := repository.New(dependency.Postgres)

	projectUseCase := projUseCase.New(repo, dependency.Logger)
	delivery.New(dependency.Router, projectUseCase, dependency.TeamUseCase)

	return projectUseCase
}
