package status

import (
	"task-tracker-server/internal/domain/status/delivery"
	"task-tracker-server/internal/domain/status/repository"
	"task-tracker-server/internal/domain/status/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
	Router   *echo.Group
}

func New(dependency Dependency) usecase.StatusUseCase {
	repo := repository.New(dependency.Postgres)

	statusUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, statusUseCase)

	return statusUseCase
}
