package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r projectRepo) DeleteProject(ctx context.Context, projectID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Delete("project").
		Where(sq.Eq{"id": projectID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("project - repository - DeleteProject - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("project - repository - DeleteProject - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("project - repository - DeleteProject - project not found")
	}

	return nil
}
