package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r statusRepo) UpdateOrders(ctx context.Context, boardID int, oldOrder, newOrder int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateOrders - r.Pool.Begin: %w", err)
	}
	defer tx.Rollback(ctx)

	if newOrder > oldOrder {
		sql, args, err := r.Builder.
			Update("status").
			Set("status_order", sq.Expr("status_order - 1")).
			Where(sq.And{
				sq.Eq{"board_id": boardID},
				sq.Gt{"status_order": oldOrder},
				sq.LtOrEq{"status_order": newOrder},
			}).
			ToSql()
		if err != nil {
			return fmt.Errorf("status - repository - UpdateOrders - r.Builder: %w", err)
		}

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("status - repository - UpdateOrders - tx.Exec: %w", err)
		}
	} else {
		sql, args, err := r.Builder.
			Update("status").
			Set("status_order", sq.Expr("status_order + 1")).
			Where(sq.And{
				sq.Eq{"board_id": boardID},
				sq.Lt{"status_order": oldOrder},
				sq.GtOrEq{"status_order": newOrder},
			}).
			ToSql()
		if err != nil {
			return fmt.Errorf("status - repository - UpdateOrders - r.Builder: %w", err)
		}

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("status - repository - UpdateOrders - tx.Exec: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("status - repository - UpdateOrders - tx.Commit: %w", err)
	}

	return nil
}
