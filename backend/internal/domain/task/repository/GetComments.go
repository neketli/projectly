package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"projectly-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetComments(ctx context.Context, taskId, lastCommentID int) ([]entity.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query := r.Builder.
		Select(
			"c.id",
			"c.text",
			"c.task_id",
			"u.id",
			"u.name",
			"u.surname",
			"u.email",
			"u.meta",
			"c.created_at",
			"c.updated_at",
		).
		From("comment c").
		Join("users u ON c.user_id = u.id").
		Where(sq.Eq{"c.task_id": taskId})

	if lastCommentID > 0 {
		query = query.Where(sq.Gt{"c.id": lastCommentID})
	}

	sqlQuery, args, err := query.
		OrderBy("c.created_at ASC").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("task - repository - GetComments - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("task - repository - GetComments - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	comments := make([]entity.Comment, 0)

	for rows.Next() {
		var comment entity.Comment
		var userMeta sql.NullString
		var createdAt sql.NullTime
		var updatedAt sql.NullTime

		if err = rows.Scan(
			&comment.ID,
			&comment.Text,
			&comment.TaskID,
			&comment.User.ID,
			&comment.User.Name,
			&comment.User.Surname,
			&comment.User.Email,
			&userMeta,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, fmt.Errorf("task - repository - GetComments - rows.Scan: %w", err)
		}

		var UserMetaJSON struct {
			AvatarURL string `json:"avatar"`
		}

		if userMeta.Valid {
			if err = json.Unmarshal([]byte(userMeta.String), &UserMetaJSON); err != nil {
				return nil, fmt.Errorf("task - repository - GetComments - json.Unmarshal(userMeta): %w", err)
			}
			comment.User.Avatar = UserMetaJSON.AvatarURL
		}

		comment.CreatedAt = createdAt.Time.Unix()
		comment.UpdatedAt = updatedAt.Time.Unix()

		comments = append(comments, comment)
	}

	return comments, nil
}
