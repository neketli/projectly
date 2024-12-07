package repository

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/user/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r userRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("user - repository - GetByEmail - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("user - repository - GetByEmail - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.User{}, fmt.Errorf("user - repository - GetByEmail - pgx.CollectRows: %w", err)
	}

	if len(users) == 0 {
		return entity.User{}, entity.ErrNoUserFound
	}

	return users[0], nil
}
