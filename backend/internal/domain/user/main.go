package user

import (
	"task-tracker-server/config"
	"task-tracker-server/internal/domain/user/repository"
	"task-tracker-server/internal/domain/user/usecase"
	"task-tracker-server/pkg/logger"
	"task-tracker-server/pkg/minio"
	"task-tracker-server/pkg/postgres"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
	S3       *minio.Minio
	Config   *config.Config
}

func New(dependency Dependency) usecase.UserUsecase {
	repo := repository.New(dependency.Postgres, dependency.S3)

	return usecase.New(repo, dependency.Logger, dependency.Config)
}
