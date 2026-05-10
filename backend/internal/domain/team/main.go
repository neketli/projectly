package team

import (
	"projectly-server/internal/domain/team/delivery"
	"projectly-server/internal/domain/team/repository"
	"projectly-server/internal/domain/team/usecase"
	userUseCase "projectly-server/internal/domain/user/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

// Dependency contains dependencies for team domain initialization.
type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres

	Router      *echo.Group
	UserUseCase userUseCase.UserUseCase
}

// New initializes the team domain with its dependencies and returns a TeamUseCase.
func New(dependency Dependency) usecase.TeamUseCase {
	repo := repository.New(dependency.Postgres)

	teamUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, teamUseCase, dependency.UserUseCase)

	return teamUseCase
}
