package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/board/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r boardRepo) GetBoardList(ctx context.Context, projectID int) ([]entity.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "title", "project_id").
		From("board").
		Where(sq.Eq{"project_id": projectID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("board - repository - GetBoardList - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("board - repository - GetBoardList - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	boards, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Board])
	if err != nil {
		return nil, fmt.Errorf("board - repository - GetBoardList - pgx.CollectRows: %w", err)
	}

	return boards, nil
}
