package status

import (
	"projectly-server/internal/domain/status/delivery"
	"projectly-server/internal/domain/status/repository"
	statusUseCase "projectly-server/internal/domain/status/usecase"
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

func New(dependency Dependency) statusUseCase.StatusUseCase {
	repo := repository.New(dependency.Postgres)

	statusUseCase := statusUseCase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, statusUseCase, dependency.TeamUseCase)

	return statusUseCase
}
