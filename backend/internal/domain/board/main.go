package board

import (
	"task-tracker-server/internal/domain/board/delivery"
	"task-tracker-server/internal/domain/board/repository"
	"task-tracker-server/internal/domain/board/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
	Router   *echo.Group
}

func New(dependency Dependency) usecase.BoardUseCase {
	repo := repository.New(dependency.Postgres)

	boardUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, boardUseCase)

	return boardUseCase
}
