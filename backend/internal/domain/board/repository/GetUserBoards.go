package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/board/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r boardRepo) GetUserBoards(ctx context.Context, userID int) ([]entity.BoardTeam, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("b.id", "b.title", "b.project_id", "p.team_id").
		From("board b").
		Join("project p ON b.project_id = p.id").
		Join("team_user tu ON tu.team_id = p.team_id").
		Where(sq.Eq{"tu.user_id": userID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("board - repository - GetUserBoards - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("board - repository - GetUserBoards - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	boards := make([]entity.BoardTeam, 4)
	for rows.Next() {
		board := entity.Board{}
		var teamID int
		if err = rows.Scan(&board.ID, &board.Title, &board.ProjectID, &teamID); err != nil {
			return nil, fmt.Errorf("board - repository - GetUserBoards - rows.Scan: %w", err)
		}

		boards = append(boards, entity.BoardTeam{
			Board:  board,
			TeamID: teamID,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("board - repository - GetUserBoards - rows.Err: %w", err)
	}

	return boards, nil
}
