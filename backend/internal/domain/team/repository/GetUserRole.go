package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) GetUserRole(ctx context.Context, teamID, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("role_id").
		From("team_roles").
		Where(sq.Eq{
			"team_id": teamID,
			"user_id": userID,
		}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("team - repository - GetUserRole - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("team - repository - GetUserRole - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var roleID int
		if err = rows.Scan(&roleID); err != nil {
			return 0, fmt.Errorf("team - repository - GetUserRole - rows.Scan: %w", err)
		}
		return roleID, nil
	} else {
		return 0, fmt.Errorf("user not in team")
	}

}
