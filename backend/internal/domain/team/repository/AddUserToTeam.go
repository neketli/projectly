package repository

import (
	"context"
	"fmt"
)

func (r teamRepo) AddUserToTeam(ctx context.Context, teamID, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("team_user").
		Columns(
			"team_id",
			"user_id",
		).
		Values(teamID, userID).
		Suffix("ON CONFLICT (team_id, user_id) DO NOTHING").
		ToSql()
	if err != nil {
		return fmt.Errorf("team - repository - AddUserToTeam - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("team - repository - AddUserToTeam - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("team - repository - AddUserToTeam - user is already in the team")
	}

	return nil
}
