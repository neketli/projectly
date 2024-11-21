package team

import (
	"task-tracker-server/internal/domain/team/repository"
	"task-tracker-server/internal/domain/team/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
}

func New(dependency Dependency) usecase.TeamUsecase {
	repo := repository.New(dependency.Postgres)

	return usecase.New(repo, dependency.Logger)
}
