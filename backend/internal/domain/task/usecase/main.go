package usecase

import (
	"context"
	"mime/multipart"
	"projectly-server/internal/domain/task/entity"
	"projectly-server/internal/domain/task/repository"
	"projectly-server/pkg/logger"
)

type TaskUseCase interface {
	CreateTask(ctx context.Context, task entity.Task) (entity.TaskDetailed, error)
	UpdateTask(ctx context.Context, task *entity.Task) (entity.TaskDetailed, error)
	UpdateTaskStatus(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetTaskList(ctx context.Context, boardID int) (map[int][]entity.Task, error)
	GetTasks(ctx context.Context, params *entity.TaskDetailedParams) ([]entity.TaskDetailed, error)

	CreateAttachment(ctx context.Context, taskId int, file *multipart.FileHeader) (string, error)
	GetAttachments(ctx context.Context, taskId int) ([]entity.Attachment, error)
	DeleteAttachment(ctx context.Context, objectName string) error
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
