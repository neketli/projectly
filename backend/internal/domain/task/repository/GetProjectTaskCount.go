package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r taskRepo) GetProjectTaskCount(ctx context.Context, projectID, statusID *int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	query := r.Builder.
		Select("COALESCE(MAX(project_index), 0)").
		From("task").
		Join("statuses s ON t.status_id = s.id").
		Join("boards b ON s.board_id = b.id")

	if projectID != nil {
		query = query.Where(sq.Eq{"b.project_id": &projectID})
	} else if statusID == nil {
		query = query.Where(sq.Eq{"s.id": &statusID})
	} else {
		return 0, fmt.Errorf("task - repository - GetProjectTaskCount - invalid arguments")
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, fmt.Errorf("task - repository - GetProjectTaskCount - r.Builder: %w", err)
	}

	var count int
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("task - repository - GetProjectTaskCount - r.Pool.QueryRow: %w", err)
	}

	return count, nil
}
