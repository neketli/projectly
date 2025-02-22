package repository

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

func (r taskRepo) CreateAttachment(ctx context.Context, reader io.Reader, filename string, taskId int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	objectName := fmt.Sprintf("attachments/%d/%s", taskId, filename)

	_, err := r.Minio.Client.PutObject(ctx, r.Minio.Bucket, objectName, reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return "", fmt.Errorf("task - repository - CreateAttachment - r.Minio.Client.PutObject: %w", err)
	}

	sql, args, err := r.Builder.
		Insert("attachment").
		Columns(
			"name",
			"task_id",
		).
		Values(
			objectName,
			taskId,
		).
		ToSql()

	if err != nil {
		return "", fmt.Errorf("task - repository - CreateAttachment - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return "", fmt.Errorf("task - repository - CreateAttachment - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return objectName, nil
}
