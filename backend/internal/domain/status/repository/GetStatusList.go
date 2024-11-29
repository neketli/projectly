package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/status/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r statusRepo) GetStatusList(ctx context.Context, boardID int) ([]entity.Status, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "board_id", "title", "order").
		From("status").
		Where(sq.Eq{"board_id": boardID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("status - repository - GetStatusList - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("status - repository - GetStatusList - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	statuses, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Status])
	if err != nil {
		return nil, fmt.Errorf("status - repository - GetStatusList - pgx.CollectRows: %w", err)
	}

	return statuses, nil
}
