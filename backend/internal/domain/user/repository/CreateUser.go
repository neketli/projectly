package repository

import (
	"context"
	"fmt"
	"task-tracker-server/internal/domain/user/entity"
)

func (r userRepo) CreateUser(ctx context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("users").
		Columns(
			"name",
			"surname",
			"email",
			"password",
		).
		Values(
			user.Name,
			user.Surname,
			user.Email,
			user.Password,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("user - repository - CreateUser - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("user - repository - CreateUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return fmt.Errorf("user - repository - CreateUser - rows.Scan: %w", err)
		}
		user.ID = id
	}

	return nil
}
