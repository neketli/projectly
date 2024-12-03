package task

import (
	"task-tracker-server/internal/domain/task/delivery"
	"task-tracker-server/internal/domain/task/repository"
	"task-tracker-server/internal/domain/task/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
	Router   *echo.Group
}

func New(dependency Dependency) usecase.TaskUseCase {
	repo := repository.New(dependency.Postgres)

	taskUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, taskUseCase)

	return taskUseCase
}
