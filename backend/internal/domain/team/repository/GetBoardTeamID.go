package repository

import (
	"context"
	"projectly-server/internal/domain/team/entity"

	"github.com/jackc/pgx/v5"
)

func (r *teamRepository) GetBoardTeamID(ctx context.Context, boardID int) (int, error) {
	query := `SELECT p.team_id FROM board b
		JOIN project p ON b.project_id = p.id
		WHERE b.id = $1`

	var teamID int
	err := r.Pool.QueryRow(ctx, query, boardID).Scan(&teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, entity.ErrBoardNotFound
		}
		return 0, err
	}
	return teamID, nil
}
