package repository

import (
	"context"
	"errors"
	"fmt"

	"projectly-server/internal/domain/board/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r boardRepo) GetBoard(ctx context.Context, boardID int) (entity.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "project_id", "title").
		From("board").
		Where(sq.Eq{"id": boardID}).
		ToSql()
	if err != nil {
		return entity.Board{}, fmt.Errorf("board - repository - GetBoard - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.Board{}, fmt.Errorf("board - repository - GetBoard - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	boards, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Board])
	if err != nil {
		return entity.Board{}, fmt.Errorf("board - repository - GetBoard - pgx.CollectRows: %w", err)
	}

	if len(boards) == 0 {
		return entity.Board{}, errors.New("no board found")
	}

	return boards[0], nil
}
