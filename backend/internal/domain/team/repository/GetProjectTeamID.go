package repository

import (
	"context"
	"projectly-server/internal/domain/team/entity"

	"github.com/jackc/pgx/v5"
)

func (r teamRepo) GetProjectTeamID(ctx context.Context, projectID int) (int, error) {
	query := `SELECT team_id FROM project WHERE id = $1`

	var teamID int
	err := r.Pool.QueryRow(ctx, query, projectID).Scan(&teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, entity.ErrProjectNotFound
		}
		return 0, err
	}
	return teamID, nil
}
