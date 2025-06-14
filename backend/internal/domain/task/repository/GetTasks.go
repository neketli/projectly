package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"projectly-server/internal/domain/task/entity"
	"projectly-server/internal/domain/task/repository/model"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetTasks(ctx context.Context, params *entity.TaskDetailedParams) ([]entity.TaskDetailed, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query := r.Builder.
		Select(
			"t.id",
			"p.code",
			"t.project_index",
			"t.title",
			"t.description",
			"t.priority",
			"t.story_points",
			"t.deadline",
			"t.created_at",
			"t.updated_at",
			"t.finished_at",
			"t.status_id",
			"t.created_user_id",
			"t.assigned_user_id",
			"s.id",
			"s.title",
			"s.hex_color",
			"au.id",
			"au.name",
			"au.surname",
			"au.email",
			"au.meta",
			"cu.id",
			"cu.name",
			"cu.surname",
			"cu.email",
			"cu.meta",
			"p.team_id",
			"p.id",
			"b.id",
		).
		From("task t").
		Join("status s ON s.id = t.status_id").
		Join("board b ON b.id = s.board_id").
		Join("project p ON p.id = b.project_id").
		LeftJoin("users cu ON cu.id = t.created_user_id").
		LeftJoin("users au ON au.id = t.assigned_user_id")

	limit := uint64(_defaultLimit)

	if params != nil {
		if params.UserID != nil && *params.UserID > 0 {
			query = query.Where(sq.Eq{"t.assigned_user_id": *params.UserID})
		}

		if params.BoardID != nil && *params.BoardID > 0 {
			query = query.Where(sq.Eq{"b.id": *params.BoardID})
		}

		if params.ProjectCode != nil && *params.ProjectCode != "" {
			query = query.Where(sq.Eq{"p.code": *params.ProjectCode})
		}

		if params.ProjectIndex != nil && *params.ProjectIndex > 0 {
			query = query.Where(sq.Eq{"t.project_index": *params.ProjectIndex})
		}

		if params.TeamID != nil && *params.TeamID > 0 {
			query = query.Where(sq.Eq{"p.team_id": *params.TeamID})
		}

		if params.Search != nil && *params.Search != "" {
			query = query.Where(sq.Or{
				sq.ILike{"t.title": fmt.Sprintf("%%%s%%", *params.Search)},
				sq.ILike{"t.description": fmt.Sprintf("%%%s%%", *params.Search)},
			})
		}

		if params.Limit != nil && *params.Limit > 0 {
			limit = *params.Limit
			query = query.Limit(limit)
		}
	}

	sqlQuery, args, err := query.OrderBy("t.updated_at DESC").ToSql()
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetUserTasks - query.ToSql: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetUserTasks - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	tasks := make([]entity.TaskDetailed, 0, limit)
	for rows.Next() {
		task := model.TaskDetailed{}

		var assignedUserMeta, createdUserMeta sql.NullString
		if err = rows.Scan(
			&task.ID,
			&task.ProjectCode,
			&task.ProjectIndex,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.StoryPoints,
			&task.Deadline,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.FinishedAt,
			&task.StatusID,
			&task.CreatedUserID,
			&task.AssignedUserID,
			&task.Status.ID,
			&task.Status.Title,
			&task.Status.HexColor,
			&task.AssignedUser.ID,
			&task.AssignedUser.Name,
			&task.AssignedUser.Surname,
			&task.AssignedUser.Email,
			&assignedUserMeta,
			&task.CreatedUser.ID,
			&task.CreatedUser.Name,
			&task.CreatedUser.Surname,
			&task.CreatedUser.Email,
			&createdUserMeta,
			&task.Meta.TeamID,
			&task.Meta.ProjectID,
			&task.Meta.BoardID,
		); err != nil {
			return nil, fmt.Errorf("task - repository - GetUserTasks - rows.Scan: %w", err)
		}

		var UserMetaJSON struct {
			AvatarURL string `json:"avatar"`
		}

		if assignedUserMeta.Valid {
			if err = json.Unmarshal([]byte(assignedUserMeta.String), &UserMetaJSON); err != nil {
				return nil, fmt.Errorf("task - repository - GetUserTasks - json.Unmarshal(assignedUserMeta): %w", err)
			}

			task.AssignedUser.Avatar = UserMetaJSON.AvatarURL
		}

		if createdUserMeta.Valid {
			if err = json.Unmarshal([]byte(createdUserMeta.String), &UserMetaJSON); err != nil {
				return nil, fmt.Errorf("task - repository - GetUserTasks - json.Unmarshal(createdUserMeta): %w", err)
			}

			task.CreatedUser.Avatar = UserMetaJSON.AvatarURL
		}

		tasks = append(tasks, task.ToEntity())
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("task - repository - GetUserTasks - rows.Err: %w", err)
	}

	return tasks, nil
}
