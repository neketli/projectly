package board

import (
	"task-tracker-server/internal/domain/board/repository"
	"task-tracker-server/internal/domain/board/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/postgres"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
}

func New(dependency Dependency) usecase.BoardUseCase {
	repo := repository.New(dependency.Postgres)

	return usecase.New(repo, dependency.Logger)
}
