package repository

import (
	"context"
	"projectly-server/internal/domain/media/entity"
	"projectly-server/pkg/minio"
)

// MediaRepository defines the interface for media data operations.
type MediaRepository interface {
	GetFile(ctx context.Context, filename string) (*entity.File, error)
}

type mediaRepo struct {
	*minio.Minio
}

// New creates a new MediaRepository instance.
func New(s3 *minio.Minio) MediaRepository {
	return &mediaRepo{s3}
}
