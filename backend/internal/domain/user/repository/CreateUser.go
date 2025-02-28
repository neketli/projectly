package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"projectly-server/internal/domain/user/entity"
)

func (r userRepo) CreateUser(ctx context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	meta, err := json.Marshal(user.Meta)
	if err != nil {
		return fmt.Errorf("user - repository - CreateUser - json.Marshal: %w", err)
	}

	sql, args, err := r.Builder.
		Insert("users").
		Columns(
			"name",
			"surname",
			"email",
			"password",
			"meta",
		).
		Values(
			user.Name,
			user.Surname,
			user.Email,
			user.Password,
			meta,
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
