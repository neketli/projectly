package usecase

import (
	"context"
	"projectly-server/internal/domain/task/entity"
	"projectly-server/internal/domain/task/repository"
	"projectly-server/pkg/logger"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	UpdateTaskStatus(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetTask(ctx context.Context, taskID int) (entity.Task, error)
	GetTaskList(ctx context.Context, boardID int) (map[int][]entity.Task, error)
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
