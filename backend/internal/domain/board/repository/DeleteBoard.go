package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r boardRepo) DeleteBoard(ctx context.Context, boardID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Delete("board").
		Where(sq.Eq{"id": boardID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("board - repository - DeleteBoard - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("board - repository - DeleteBoard - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("board - repository - DeleteBoard - board not found")
	}

	return nil
}
