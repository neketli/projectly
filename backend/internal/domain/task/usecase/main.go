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

	CreateAttachment(ctx context.Context, taskID int, file *multipart.FileHeader) (string, error)
	GetAttachments(ctx context.Context, taskID int) ([]entity.Attachment, error)
	DeleteAttachment(ctx context.Context, objectName string) error

	CreateComment(ctx context.Context, taskID, userID int, text string) error
	GetComments(ctx context.Context, taskID, lastCommentID int) ([]entity.Comment, error)
	DeleteComment(ctx context.Context, taskID, commentID int) error
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
