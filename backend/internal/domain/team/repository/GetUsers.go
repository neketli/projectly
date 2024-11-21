package repository

import (
	"context"
	"fmt"

	userEntity "task-tracker-server/internal/domain/user/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r teamRepo) GetUsers(ctx context.Context, teamID int) ([]userEntity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select(
			"tu.user_id",
			"u.email",
			"u.name",
			"u.surname",
			"u.meta",
		).
		From("team_users tu").
		Join("users u on u.id = tu.user_id").
		Where(sq.Eq{"tu.team_id": teamID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetUsers - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetUsers - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[userEntity.User])
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetUsers - pgx.CollectRows: %w", err)
	}

	return users, nil
}
