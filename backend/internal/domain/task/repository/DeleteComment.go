package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) DeleteComment(ctx context.Context, taskID, commentID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Delete("comment").
		Where(sq.Eq{"id": commentID, "task_id": taskID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("task - repository - DeleteComment - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("task - repository - DeleteComment - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
