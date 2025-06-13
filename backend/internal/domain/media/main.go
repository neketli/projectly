package media

import (
	"projectly-server/internal/domain/media/delivery"
	"projectly-server/internal/domain/media/repository"
	"projectly-server/internal/domain/media/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/minio"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger *logger.Logger
	S3     *minio.Minio
	Router *echo.Group
}

func New(dependency Dependency) usecase.MediaUseCase {
	repo := repository.New(dependency.S3)

	mediaUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, mediaUseCase)

	return mediaUseCase
}
