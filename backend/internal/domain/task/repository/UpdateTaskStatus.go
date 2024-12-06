package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) UpdateTaskStatus(ctx context.Context, task *entity.Task) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query, args, err := r.Builder.
		Update("task").
		SetMap(sq.Eq{
			"status_id": task.StatusID,
		}).
		Where(sq.Eq{"id": task.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("task - repository - UpdateTaskStatus - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("task - repository - UpdateTaskStatus - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task - repository - UpdateTaskStatus - task not found")
	}

	return nil
}
