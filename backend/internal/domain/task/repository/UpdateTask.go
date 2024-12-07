package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"projectly-server/internal/domain/task/entity"
	"projectly-server/internal/domain/task/repository/model"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) UpdateTask(ctx context.Context, task *entity.Task) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	taskModel := model.Task{
		Title:          task.Title,
		Description:    sql.NullString{String: task.Description, Valid: true},
		Priority:       sql.NullInt64{Int64: int64(task.Priority), Valid: task.Priority > 0},
		StoryPoints:    sql.NullInt64{Int64: int64(task.StoryPoints), Valid: task.StoryPoints > 0},
		TrackedTime:    sql.NullInt64{Int64: int64(task.TrackedTime), Valid: task.TrackedTime > 0},
		StatusID:       task.StatusID,
		UpdatedAt:      sql.NullTime{Time: time.Now().UTC(), Valid: true},
		Deadline:       sql.NullTime{Time: time.Unix(task.Deadline, 0).UTC(), Valid: task.Deadline != 0},
		FinishedAt:     sql.NullTime{Time: time.Unix(task.FinishedAt, 0).UTC(), Valid: task.FinishedAt != 0},
		AssignedUserID: sql.NullInt64{Int64: int64(task.AssignedUserID), Valid: task.AssignedUserID != 0},
	}

	query, args, err := r.Builder.
		Update("task").
		SetMap(sq.Eq{
			"title":            taskModel.Title,
			"description":      taskModel.Description,
			"priority":         taskModel.Priority,
			"story_points":     taskModel.StoryPoints,
			"tracked_time":     taskModel.TrackedTime,
			"deadline":         taskModel.Deadline,
			"updated_at":       taskModel.UpdatedAt,
			"finished_at":      taskModel.FinishedAt,
			"status_id":        taskModel.StatusID,
			"assigned_user_id": taskModel.AssignedUserID,
		}).
		Where(sq.Eq{"id": task.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("task - repository - UpdateTask - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("task - repository - UpdateTask - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task - repository - UpdateTask - task not found")
	}

	return nil
}
