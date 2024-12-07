package user

import (
	"projectly-server/config"
	"projectly-server/internal/domain/user/delivery"
	"projectly-server/internal/domain/user/repository"
	"projectly-server/internal/domain/user/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/minio"
	"projectly-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger     *logger.Logger
	Postgres   *postgres.Postgres
	S3         *minio.Minio
	Config     *config.Config
	Router     *echo.Group
	AuthRouter *echo.Group
}

func New(dependency Dependency) usecase.UserUseCase {
	repo := repository.New(dependency.Postgres, dependency.S3)
	userUseCase := usecase.New(repo, dependency.Logger, dependency.Config)

	delivery.New(dependency.AuthRouter, dependency.Router, userUseCase)

	return userUseCase
}
