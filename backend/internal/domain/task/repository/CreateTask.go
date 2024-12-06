package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"task-tracker-server/internal/domain/task/entity"
	"task-tracker-server/internal/domain/task/repository/model"
)

func (r taskRepo) CreateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	projectID, err := r.getProjectIdByStatus(ctx, task.StatusID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.getProjectIdByStatus: %w", err)
	}

	projectTaskIndex, err := r.getProjectTaskIndex(ctx, projectID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.GetProjectTaskIndex: %w", err)
	}

	assigned := sql.NullInt64{Int64: 0, Valid: false}
	if task.AssignedUserID == 0 {
		assigned = sql.NullInt64{Int64: int64(task.AssignedUserID), Valid: false}
	}

	finishedAt := sql.NullTime{Time: time.Time{}, Valid: false}
	if task.FinishedAt != 0 {
		finishedAt = sql.NullTime{Time: time.Unix(task.FinishedAt, 0).UTC(), Valid: true}
	}

	query, args, err := r.Builder.
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
			"finished_at",
		).
		Values(
			projectTaskIndex+1,
			task.Title,
			task.Description,
			task.Priority,
			task.StoryPoints,
			task.TrackedTime,
			time.Unix(task.Deadline, 0).UTC(),
			task.StatusID,
			task.CreatedUserID,
			assigned,
			finishedAt,
		).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return entity.Task{}, fmt.Errorf("task - repository - CreateTask - failed to create task")
	}

	var result model.Task
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

	return result.ToEntity(), nil
}
