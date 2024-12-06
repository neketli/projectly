package repository

import (
	"context"
	"task-tracker-server/internal/domain/task/entity"
	"task-tracker-server/pkg/postgres"
	"time"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	UpdateTaskStatus(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetTask(ctx context.Context, taskID int) (entity.Task, error)
	GetTaskList(ctx context.Context, boardID int, limit uint64) ([]entity.Task, error)
	GetUserTasks(ctx context.Context, userID int, limit uint64) ([]entity.TaskCard, error)
}

const (
	_defaultConnTimeout = 5 * time.Second
	_defaultLimit       = 10
)

type taskRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) TaskRepository {
	return taskRepo{pg}
}
