package board

import (
	"projectly-server/internal/domain/board/delivery"
	"projectly-server/internal/domain/board/repository"
	"projectly-server/internal/domain/board/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/postgres"

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
