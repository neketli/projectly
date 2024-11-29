package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r statusRepo) DeleteStatus(ctx context.Context, statusID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Delete("status").
		Where(sq.Eq{"id": statusID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("status - repository - DeleteStatus - status not found")
	}

	return nil
}
