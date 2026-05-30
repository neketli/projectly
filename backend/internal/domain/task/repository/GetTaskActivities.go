package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"projectly-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetTaskActivities(ctx context.Context, taskID int) ([]entity.TaskActivity, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query, args, err := r.Builder.
		Select(
			"a.id",
			"a.task_id",
			"u.id",
			"u.name",
			"u.surname",
			"u.email",
			"u.meta",
			"a.action_type",
			"a.field_name",
			"a.old_value",
			"a.new_value",
			"a.created_at",
		).
		From("task_activity a").
		Join("users u ON a.user_id = u.id").
		Where(sq.Eq{"a.task_id": taskID}).
		OrderBy("a.created_at DESC").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskActivities - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetTaskActivities - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	activities := make([]entity.TaskActivity, 0)

	for rows.Next() {
		var activity entity.TaskActivity
		var userMeta sql.NullString
		var createdAt sql.NullTime
		var fieldName sql.NullString
		var oldValue sql.NullString
		var newValue sql.NullString

		if err = rows.Scan(
			&activity.ID,
			&activity.TaskID,
			&activity.User.ID,
			&activity.User.Name,
			&activity.User.Surname,
			&activity.User.Email,
			&userMeta,
			&activity.ActionType,
			&fieldName,
			&oldValue,
			&newValue,
			&createdAt,
		); err != nil {
			return nil, fmt.Errorf("task - repository - GetTaskActivities - rows.Scan: %w", err)
		}

		var userMetaJSON struct {
			AvatarURL string `json:"avatar"`
		}
		if userMeta.Valid {
			if err = json.Unmarshal([]byte(userMeta.String), &userMetaJSON); err != nil {
				return nil, fmt.Errorf("task - repository - GetTaskActivities - json.Unmarshal(userMeta): %w", err)
			}
			activity.User.Avatar = userMetaJSON.AvatarURL
		}

		activity.FieldName = fieldName.String
		activity.OldValue = oldValue.String
		activity.NewValue = newValue.String
		activity.CreatedAt = createdAt.Time.Unix()

		activities = append(activities, activity)
	}

	return activities, nil
}
