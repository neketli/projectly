package team

import (
	"task-tracker-server/internal/domain/team/delivery"
	"task-tracker-server/internal/domain/team/repository"
	"task-tracker-server/internal/domain/team/usecase"
	userUseCase "task-tracker-server/internal/domain/user/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"

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
