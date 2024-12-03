package usecase

import (
	"context"
	"task-tracker-server/internal/domain/task/entity"
	"task-tracker-server/internal/domain/task/repository"
	"task-tracker-server/pkg/logger"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetTask(ctx context.Context, taskID int) (entity.Task, error)
	GetTaskList(ctx context.Context, boardID int, limit uint64) ([]entity.Task, error)
	GetUserTasks(ctx context.Context, userID int, limit uint64) ([]entity.TaskCard, error)
}

type taskUseCase struct {
	repo   repository.TaskRepository
	logger *logger.Logger
}

func New(r repository.TaskRepository, l *logger.Logger) TaskUseCase {
	return &taskUseCase{
		repo:   r,
		logger: l,
	}
}
