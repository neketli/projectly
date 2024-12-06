package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) getProjectIdByStatus(ctx context.Context, statusID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("b.project_id").
		From("board b").
		Join("status s ON s.board_id = b.id").
		Where(sq.Eq{"s.id": statusID}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("task - repository - getProjectIdByStatus - r.Builder: %w", err)
	}

	var projectID int
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&projectID)
	if err != nil {
		return 0, fmt.Errorf("task - repository - getProjectIdByStatus - r.Pool.QueryRow: %w", err)
	}

	return projectID, nil
}
