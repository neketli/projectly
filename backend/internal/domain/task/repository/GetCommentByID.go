package repository

import (
	"context"
	"database/sql"
	"fmt"
	"projectly-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetCommentByID(ctx context.Context, commentID int) (entity.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query, args, err := r.Builder.
		Select(
			"id",
			"text",
			"task_id",
			"user_id",
			"created_at",
			"updated_at",
		).
		From("comment").
		Where(sq.Eq{"id": commentID}).
		ToSql()

	if err != nil {
		return entity.Comment{}, fmt.Errorf("task - repository - GetCommentByID - r.Builder: %w", err)
	}

	var comment entity.Comment
	var createdAt sql.NullTime
	var updatedAt sql.NullTime
	var userID int

	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&comment.ID,
		&comment.Text,
		&comment.TaskID,
		&userID,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return entity.Comment{}, fmt.Errorf("task - repository - GetCommentByID - r.Pool.QueryRow: %w", err)
	}

	comment.User.ID = userID
	comment.CreatedAt = createdAt.Time.Unix()
	comment.UpdatedAt = updatedAt.Time.Unix()

	return comment, nil
}
