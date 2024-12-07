package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/board/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r boardRepo) UpdateBoard(ctx context.Context, board *entity.Board) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Update("board").
		SetMap(sq.Eq{
			"title": board.Title,
		}).
		Where(sq.Eq{"id": board.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("board - repository - UpdateBoard - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("board - repository - UpdateBoard - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("board - repository - UpdateBoard - board not found")
	}

	return nil
}
