package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetTaskList(ctx context.Context, boardID int, limit uint64) ([]entity.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	if limit <= 0 {
		limit = _defaultLimit
	}

	sql, args, err := r.Builder.
		Select(
			"id",
			"project_index",
			"title",
			"description",
			"priority",
			"story_points",
			"tracked_time",
			"deadline",
			"created_at",
			"updated_at",
			"finished_at",
			"status_id",
			"created_user_id",
			"assigned_user_id",
		).
		From("task").
		Where(sq.Eq{"board_id": boardID}).
		OrderBy("updated_at ASC").
		Limit(limit).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskList - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskList - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	tasks := make([]entity.Task, 0)
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(
			&task.ID,
			&task.ProjectIndex,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.StoryPoints,
			&task.TrackedTime,
			&task.Deadline,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.FinishedAt,
			&task.StatusID,
			&task.CreatedUserID,
			&task.AssignedUserID,
		); err != nil {
			return nil, fmt.Errorf("task - repository - GetTaskList - rows.Scan: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskList - rows.Err: %w", err)
	}

	return tasks, nil
}
