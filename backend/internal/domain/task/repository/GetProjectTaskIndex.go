package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) getProjectTaskIndex(ctx context.Context, projectID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("COALESCE(MAX(t.project_index), 0)").
		From("task t").
		Join("status s ON t.status_id = s.id").
		Join("board b ON s.board_id = b.id").
		Where(sq.Eq{"b.project_id": &projectID}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("task - repository - GetProjectTaskIndex - r.Builder: %w", err)
	}

	var idx int
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&idx)
	if err != nil {
		return 0, fmt.Errorf("task - repository - GetProjectTaskIndex - r.Pool.QueryRow: %w", err)
	}

	return idx, nil
}
