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

type Dependency struct {
	Logger       *logger.Logger
	Postgres     *postgres.Postgres
	Router       *echo.Group
	TeamUseCase  teamUseCase.TeamUseCase
}

func New(dependency Dependency) projUseCase.ProjectUseCase {
	repo := repository.New(dependency.Postgres)

	projectUseCase := projUseCase.New(repo, dependency.Logger)
	delivery.New(dependency.Router, projectUseCase, dependency.TeamUseCase)

	return projectUseCase
}
