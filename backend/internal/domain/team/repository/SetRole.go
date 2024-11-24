package repository

import (
	"context"
	"fmt"
)

func (r teamRepo) SetRole(ctx context.Context, teamID, userID, roleID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("team_roles").
		Columns("team_id", "user_id", "role_id").
		Values(teamID, userID, roleID).
		Suffix("ON CONFLICT (team_id, user_id) DO UPDATE SET role_id = EXCLUDED.role_id").
		ToSql()
	if err != nil {
		return fmt.Errorf("team - repository - SetRoles - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("team - repository - SetRoles - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("team - repository - SetRoles - user is already in the team")
	}

	return nil
}
