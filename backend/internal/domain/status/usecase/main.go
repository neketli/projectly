package usecase

import (
	"context"
	"task-tracker-server/internal/domain/status/entity"
	"task-tracker-server/internal/domain/status/repository"
	"task-tracker-server/pkg/logger"
)

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

func New(r repository.StatusRepository, l *logger.Logger) StatusUseCase {
	return &statusUseCase{
		repo:   r,
		logger: l,
	}
}
