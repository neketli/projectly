package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r taskRepo) CreateTaskActivity(ctx context.Context, taskID, userID int, actionType, fieldName, oldValue, newValue string) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query, args, err := r.Builder.
		Insert("task_activity").
		Columns(
			"task_id",
			"user_id",
			"action_type",
			"field_name",
			"old_value",
			"new_value",
			"created_at",
		).
		Values(
			taskID,
			userID,
			actionType,
			sql.NullString{String: fieldName, Valid: fieldName != ""},
			sql.NullString{String: oldValue, Valid: oldValue != ""},
			sql.NullString{String: newValue, Valid: newValue != ""},
			sql.NullTime{Time: time.Now().UTC(), Valid: true},
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("task - repository - CreateTaskActivity - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("task - repository - CreateTaskActivity - r.Pool.Exec: %w", err)
	}

	return nil
}
