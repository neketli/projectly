package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r statusRepo) DeleteStatus(ctx context.Context, statusID, order int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - r.Pool.Begin: %w", err)
	}
	defer tx.Rollback(ctx)

	// Delete status
	sql, args, err := r.Builder.
		Delete("status").
		Where(sq.Eq{"id": statusID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - r.Builder: %w", err)
	}

	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - tx.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("status - repository - DeleteStatus - status not found")
	}

	// Update orders in board
	sql, args, err = r.Builder.
		Update("status").
		Set("status_order", sq.Expr("status_order - 1")).
		Where(sq.Gt{"status_order": order}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - r.Builder: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - tx.Exec: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("status - repository - DeleteStatus - tx.Commit: %w", err)
	}

	return nil
}
