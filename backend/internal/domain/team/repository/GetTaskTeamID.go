package repository

import (
	"context"
	"projectly-server/internal/domain/team/entity"

	"github.com/jackc/pgx/v5"
)

func (r teamRepo) GetTaskTeamID(ctx context.Context, taskID int) (int, error) {
	query := `SELECT p.team_id FROM task t
		JOIN status s ON t.status_id = s.id
		JOIN board b ON s.board_id = b.id
		JOIN project p ON b.project_id = p.id
		WHERE t.id = $1`

	var teamID int
	err := r.Pool.QueryRow(ctx, query, taskID).Scan(&teamID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, entity.ErrTaskNotFound
		}
		return 0, err
	}
	return teamID, nil
}
