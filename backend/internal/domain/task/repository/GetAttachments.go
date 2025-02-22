package repository

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetAttachments(ctx context.Context, taskId int) ([]entity.Attachment, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select(
			"id",
			"name",
			"task_id",
		).
		From("attachment").
		Where(sq.Eq{"task_id": taskId}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("task - repository - CreateAttachment - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - CreateAttachment - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var attachments []entity.Attachment

	for rows.Next() {
		var attachment entity.Attachment
		if err = rows.Scan(&attachment.ID, &attachment.Name, &attachment.TaskID); err != nil {
			return nil, fmt.Errorf("task - repository - CreateAttachment - rows.Scan: %w", err)
		}

		attachments = append(attachments, attachment)
	}

	return attachments, nil
}
