package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/status/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r statusRepo) CreateStatus(ctx context.Context, status *entity.Status) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Update("status").
		Set("status_order", sq.Expr("status_order + 1")).
		Where(sq.And{
			sq.Eq{"board_id": status.BoardID},
			sq.GtOrEq{"status_order": status.Order},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - UpdateOrder - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateOrder - r.Pool.Exec: %w", err)
	}

	sql, args, err = r.Builder.
		Insert("status").
		Columns(
			"title",
			"board_id",
			"status_order",
			"hex_color",
		).
		Values(
			status.Title,
			status.BoardID,
			status.Order,
			status.HexColor,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - CreateStatus - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - CreateStatus - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("status - repository - CreateStatus - status not found")
	}

	var id int
	if err = rows.Scan(&id); err != nil {
		return fmt.Errorf("status - repository - CreateStatus - rows.Scan: %w", err)
	}

	status.ID = id

	return nil
}
