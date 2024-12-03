package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) UpdateTask(ctx context.Context, task *entity.Task) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Update("task").
		SetMap(sq.Eq{
			"title":            task.Title,
			"description":      task.Description,
			"priority":         task.Priority,
			"story_points":     task.StoryPoints,
			"tracked_time":     task.TrackedTime,
			"deadline":         task.Deadline,
			"status_id":        task.StatusID,
			"assigned_user_id": task.AssignedUserID,
		}).
		Where(sq.Eq{"id": task.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("task - repository - UpdateTask - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("task - repository - UpdateTask - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task - repository - UpdateTask - task not found")
	}

	return nil
}
