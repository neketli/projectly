package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r taskRepo) CreateComment(ctx context.Context, taskID, userID int, text string) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("comment").
		Columns(
			"text",
			"task_id",
			"user_id",
			"created_at",
			"updated_at",
		).
		Values(
			text,
			taskID,
			userID,
			sql.NullTime{Time: time.Now().UTC(), Valid: true},
			sql.NullTime{Time: time.Now().UTC(), Valid: true},
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("task - repository - CreateComment - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("task - repository - CreateComment - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
