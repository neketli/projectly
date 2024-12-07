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

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres

	Router      *echo.Group
	UserUseCase userUseCase.UserUseCase
}

func New(dependency Dependency) usecase.TeamUseCase {
	repo := repository.New(dependency.Postgres)

	teamUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, teamUseCase, dependency.UserUseCase)

	return teamUseCase
}
