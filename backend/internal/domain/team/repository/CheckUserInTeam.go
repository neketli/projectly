package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) CheckUserInTeam(ctx context.Context, teamID, userID int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("user_id").
		From("team_user").
		Where(sq.Eq{
			"team_id": teamID,
			"user_id": userID,
		}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("team - repository - CheckUserInTeam - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("team - repository - CheckUserInTeam - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
