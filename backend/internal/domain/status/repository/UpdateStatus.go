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

	sql, args, err := r.Builder.
		Update("status").
		SetMap(sq.Eq{
			"title": status.Title,
			"order": status.Order,
		}).
		Where(sq.Eq{"id": status.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("status - repository - UpdateStatus - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("status - repository - UpdateStatus - status not found")
	}

	return nil
}
