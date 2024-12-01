package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/status/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r statusRepo) UpdateStatus(ctx context.Context, status *entity.Status) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - r.Pool.Begin: %w", err)
	}
	defer tx.Rollback(ctx)

	// update orders in board
	sql, args, err := r.Builder.
		Update("status").
		Set("status_order", sq.Expr("status_order + 1")).
		Where(sq.And{
			sq.Eq{"board_id": status.BoardID},
			sq.GtOrEq{"status_order": status.Order},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - r.Builder: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - tx.Exec: %w", err)
	}

	// update status
	sql, args, err = r.Builder.
		Update("status").
		SetMap(sq.Eq{
			"title":        status.Title,
			"status_order": status.Order,
			"hex_color":    status.HexColor,
		}).
		Where(sq.Eq{"id": status.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - r.Builder: %w", err)
	}

	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - tx.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("status - repository - UpdateStatus - status not found")
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - tx.Commit: %w", err)
	}

	return nil
}
