package repository

import (
	"context"
	"fmt"
	"task-tracker-server/internal/domain/user/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Update("users").
		SetMap(sq.Eq{
			"name":     user.Name,
			"surname":  user.Surname,
			"email":    user.Email,
			"password": user.Password,
		}).
		Where(sq.Eq{"id": user.ID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("user - repository - UpdateUser - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("user - repository - UpdateUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
