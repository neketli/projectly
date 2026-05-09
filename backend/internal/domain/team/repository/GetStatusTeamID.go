package repository

import (
	"context"
	"projectly-server/internal/domain/team/entity"

	"github.com/jackc/pgx/v5"
)

func (r *teamRepository) GetStatusTeamID(ctx context.Context, statusID int) (int, error) {
	query := `SELECT p.team_id FROM status s
		JOIN board b ON s.board_id = b.id
		JOIN project p ON b.project_id = p.id
		WHERE s.id = $1`

	var teamID int
	err := r.Pool.QueryRow(ctx, query, statusID).Scan(&teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, entity.ErrStatusNotFound
		}
		return 0, err
	}
	return teamID, nil
}
