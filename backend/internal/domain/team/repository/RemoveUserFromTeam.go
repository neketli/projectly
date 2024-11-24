package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) RemoveUserFromTeam(ctx context.Context, teamID, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Delete("team_user").
		Where(sq.Eq{
			"team_id": teamID,
			"user_id": userID,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("team - repository - RemoveUserFromTeam - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("team - repository - RemoveUserFromTeam - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("team - repository - RemoveUserFromTeam - user not found in the team")
	}

	return nil
}
