package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateOrders updates status orders when reordering.
func (r statusRepo) UpdateOrders(ctx context.Context, boardID, oldOrder, newOrder int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateOrders - r.Pool.Begin: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx) //nolint:errcheck //rollback result intentionally ignored
	}()

	if newOrder > oldOrder {
		query, queryArgs, queryErr := r.Builder.
			Update("status").
			Set("status_order", sq.Expr("status_order - 1")).
			Where(sq.And{
				sq.Eq{"board_id": boardID},
				sq.Gt{"status_order": oldOrder},
				sq.LtOrEq{"status_order": newOrder},
			}).
			ToSql()
		if queryErr != nil {
			return fmt.Errorf("status - repository - UpdateOrders - r.Builder: %w", queryErr)
		}

		_, err = tx.Exec(ctx, query, queryArgs...)
		if err != nil {
			return fmt.Errorf("status - repository - UpdateOrders - tx.Exec: %w", err)
		}
	} else {
		query, queryArgs, queryErr := r.Builder.
			Update("status").
			Set("status_order", sq.Expr("status_order + 1")).
			Where(sq.And{
				sq.Eq{"board_id": boardID},
				sq.Lt{"status_order": oldOrder},
				sq.GtOrEq{"status_order": newOrder},
			}).
			ToSql()
		if queryErr != nil {
			return fmt.Errorf("status - repository - UpdateOrders - r.Builder: %w", queryErr)
		}

		_, err = tx.Exec(ctx, query, queryArgs...)
		if err != nil {
			return fmt.Errorf("status - repository - UpdateOrders - tx.Exec: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("status - repository - UpdateOrders - tx.Commit: %w", err)
	}

	return nil
}
