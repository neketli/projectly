package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"
	"task-tracker-server/internal/domain/task/repository/model"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetTaskList(ctx context.Context, boardID int, limit uint64) ([]entity.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	if limit <= 0 {
		limit = _defaultLimit
	}

	query, args, err := r.Builder.
		Select(
			"t.id",
			"t.project_index",
			"t.title",
			"t.description",
			"t.priority",
			"t.story_points",
			"t.tracked_time",
			"t.deadline",
			"t.created_at",
			"t.updated_at",
			"t.finished_at",
			"t.status_id",
			"t.created_user_id",
			"t.assigned_user_id",
		).
		From("task t").
		Join("status s ON t.status_id = s.id").
		Join("board b ON s.board_id = b.id").
		Where(sq.Eq{"board_id": boardID}).
		OrderBy("updated_at ASC").
		Limit(limit).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskList - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskList - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	tasks := make([]entity.Task, 0)
	for rows.Next() {
		var task model.Task
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

		tasks = append(tasks, task.ToEntity())
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskList - rows.Err: %w", err)
	}

	return tasks, nil
}
