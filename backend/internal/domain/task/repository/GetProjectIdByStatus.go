package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) getProjectIDByStatus(ctx context.Context, statusID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query, args, err := r.Builder.
		Select("b.project_id").
		From("board b").
		Join("status s ON s.board_id = b.id").
		Where(sq.Eq{"s.id": statusID}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("task - repository - getProjectIDByStatus - r.Builder: %w", err)
	}

	var projectID int
	err = r.Pool.QueryRow(ctx, query, args...).Scan(&projectID)
	if err != nil {
		return 0, fmt.Errorf("task - repository - getProjectIDByStatus - r.Pool.QueryRow: %w", err)
	}

	return projectID, nil
}
