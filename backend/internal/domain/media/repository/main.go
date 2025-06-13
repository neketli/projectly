package repository

import (
	"context"
	"projectly-server/internal/domain/media/entity"
	"projectly-server/pkg/minio"
)

type MediaRepository interface {
	GetFile(ctx context.Context, filename string) (*entity.File, error)
}

type mediaRepo struct {
	*minio.Minio
}

func New(s3 *minio.Minio) MediaRepository {
	return &mediaRepo{s3}
}
