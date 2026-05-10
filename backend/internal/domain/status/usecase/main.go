package usecase

import (
	"context"
	"projectly-server/internal/domain/status/entity"
	"projectly-server/internal/domain/status/repository"
	"projectly-server/pkg/logger"
)

// StatusUseCase defines the interface for status business logic.
type StatusUseCase interface {
	CreateStatus(ctx context.Context, status *entity.Status) error
	UpdateStatus(ctx context.Context, status *entity.Status, oldOrder *int) error
	DeleteStatus(ctx context.Context, statusID, order int) error
	GetStatusList(ctx context.Context, boardID int) ([]entity.Status, error)
}

type statusUseCase struct {
	repo   repository.StatusRepository
	logger *logger.Logger
}

// New creates a new StatusUseCase instance.
func New(r repository.StatusRepository, l *logger.Logger) StatusUseCase {
	return &statusUseCase{
		repo:   r,
		logger: l,
	}
}
