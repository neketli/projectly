package repository

import (
	"context"
	"projectly-server/internal/domain/task/entity"
	"projectly-server/pkg/postgres"
	"time"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	UpdateTaskStatus(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetTaskList(ctx context.Context, boardID int) ([]entity.Task, error)
	GetTasks(ctx context.Context, params *entity.TaskDetailedParams) ([]entity.TaskDetailed, error)
}

const (
	_defaultConnTimeout = 30 * time.Second
	_defaultLimit       = 10
)

type taskRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) TaskRepository {
	return taskRepo{pg}
}
