package board

import (
	"projectly-server/internal/domain/board/delivery"
	"projectly-server/internal/domain/board/repository"
	boardUseCase "projectly-server/internal/domain/board/usecase"
	teamUseCase "projectly-server/internal/domain/team/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger      *logger.Logger
	Postgres    *postgres.Postgres
	Router      *echo.Group
	TeamUseCase teamUseCase.TeamUseCase
}

func New(dependency Dependency) boardUseCase.BoardUseCase {
	repo := repository.New(dependency.Postgres)

	boardUseCase := boardUseCase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, boardUseCase, dependency.TeamUseCase)

	return boardUseCase
}
