package project

import (
	"task-tracker-server/internal/domain/project/repository"
	"task-tracker-server/internal/domain/project/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
}

func New(dependency Dependency) usecase.ProjectUseCase {
	repo := repository.New(dependency.Postgres)

	return usecase.New(repo, dependency.Logger)
}
