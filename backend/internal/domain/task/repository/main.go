package repository

import (
	"context"
	"io"
	"projectly-server/internal/domain/task/entity"
	"projectly-server/pkg/minio"
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

	CreateAttachment(ctx context.Context, reader io.Reader, filename string, taskID int) (string, error)
	GetAttachments(ctx context.Context, taskID int) ([]entity.Attachment, error)
	DeleteAttachment(ctx context.Context, objectName string) error

	CreateComment(ctx context.Context, taskID, userID int, text string) error
	GetComments(ctx context.Context, taskID, lastCommentID int) ([]entity.Comment, error)
	DeleteComment(ctx context.Context, taskID, commentID int) error
}

const (
	_defaultConnTimeout = 30 * time.Second
	_defaultLimit       = 10
)

type taskRepo struct {
	*postgres.Postgres
	*minio.Minio
}

func New(pg *postgres.Postgres, s3 *minio.Minio) TaskRepository {
	return taskRepo{pg, s3}
}
