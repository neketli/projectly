package usecase

import (
	"context"
	"projectly-server/internal/domain/media/entity"
	"projectly-server/internal/domain/media/repository"
	"projectly-server/pkg/logger"
)

// MediaUseCase defines the interface for media business logic.
type MediaUseCase interface {
	GetFile(ctx context.Context, filename string) (*entity.File, error)
}

type mediaUseCase struct {
	repo   repository.MediaRepository
	logger *logger.Logger
}

// New creates a new MediaUseCase instance.
func New(r repository.MediaRepository, l *logger.Logger) MediaUseCase {
	return &mediaUseCase{
		repo:   r,
		logger: l,
	}
}
