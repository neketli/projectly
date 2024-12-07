package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/status/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r statusRepo) GetStatusList(ctx context.Context, boardID int) ([]entity.Status, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "board_id", "title", "status_order", "hex_color").
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

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("status - repository - GetStatusList - rows.Err: %w", err)
	}

	statuses := make([]entity.Status, 0, 3)
	for rows.Next() {
		var status entity.Status
		if err := rows.Scan(
			&status.ID,
			&status.BoardID,
			&status.Title,
			&status.Order,
			&status.HexColor,
		); err != nil {
			return nil, fmt.Errorf("status - repository - GetStatusList - rows.Scan: %w", err)
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}
