package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"
	userEntity "task-tracker-server/internal/domain/user/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetUserTasks(ctx context.Context, userID int, limit uint64) ([]entity.TaskCard, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	if limit <= 0 {
		limit = _defaultLimit
	}

	sql, args, err := r.Builder.
		Select(
			"p.code",
			"t.project_index",
			"t.title",
			"t.priority",
			"t.story_points",
			"t.deadline",
			"s.title",
			"s.hex_color",
			"p.team_id",
			"u.name",
			"u.surname",
			"u.meta",
		).
		From("task t").
		Join("status s ON s.id = t.status_id").
		Join("board b ON b.id = s.board_id").
		Join("project p ON p.id = b.project_id").
		Join("team_user tu ON tu.team_id = p.team_id").
		Join("user u ON u.id = tu.team_id").
		Where(sq.Eq{"tu.user_id": userID}).
		OrderBy("t.updated_at DESC").
		Limit(limit).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetUserTasks - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetUserTasks - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	tasks := make([]entity.TaskCard, 0, limit)
	for rows.Next() {
		task := entity.TaskCard{}
		userMeta := userEntity.UserMeta{}
		if err = rows.Scan(
			&task.ProjectCode,
			&task.ProjectIndex,
			&task.Title,
			&task.Priority,
			&task.StoryPoints,
			&task.Deadline,
			&task.Status.Title,
			&task.Status.HexColor,
			&task.AssignedUser.Name,
			&task.AssignedUser.Surname,
			&userMeta,
		); err != nil {
			return nil, fmt.Errorf("task - repository - GetUserTasks - rows.Scan: %w", err)
		}

		task.AssignedUser.Avatar = userMeta.Avatar

		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("task - repository - GetUserTasks - rows.Err: %w", err)
	}

	return tasks, nil
}
