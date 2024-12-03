package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"
)

func (r taskRepo) CreateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	ProjectTaskCount, err := r.GetProjectTaskCount(ctx, nil, &task.StatusID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.GetProjectTaskCount: %w", err)
	}

	sql, args, err := r.Builder.
		Insert("task").
		Columns(
			"project_index",
			"title",
			"description",
			"priority",
			"story_points",
			"tracked_time",
			"deadline",
			"status_id",
			"created_user_id",
			"assigned_user_id",
		).
		Values(
			ProjectTaskCount+1,
			task.Title,
			task.Description,
			task.Priority,
			task.StoryPoints,
			task.TrackedTime,
			task.Deadline,
			task.StatusID,
			task.CreatedUserID,
			task.AssignedUserID,
		).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - task not found")
	}

	var result entity.Task
	if err = rows.Scan(
		&result.ID,
		&result.ProjectIndex,
		&result.Title,
		&result.Description,
		&result.Priority,
		&result.StoryPoints,
		&result.TrackedTime,
		&result.Deadline,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.FinishedAt,
		&result.StatusID,
		&result.CreatedUserID,
		&result.AssignedUserID,
	); err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - rows.Scan: %w", err)
	}
	return result, nil
}
