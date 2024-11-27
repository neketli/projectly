package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/board/entity"
)

func (r boardRepo) CreateBoard(ctx context.Context, board *entity.Board) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("board").
		Columns(
			"title",
			"project_id",
		).
		Values(
			board.Title,
			board.ProjectID,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("board - repository - CreateBoard - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("board - repository - CreateBoard - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("board - repository - CreateBoard - board not found")
	}

	var id int
	if err = rows.Scan(&id); err != nil {
		return fmt.Errorf("board - repository - CreateBoard - rows.Scan: %w", err)
	}

	board.ID = id
	return nil
}
