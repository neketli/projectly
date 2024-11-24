package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/team/entity"

	"github.com/jackc/pgx/v5"
)

func (r teamRepo) GetRoles(ctx context.Context) ([]entity.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "role_name").
		From("roles").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetRoles - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetRoles - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Role])
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetRoles - pgx.CollectRows: %w", err)
	}

	return roles, nil
}
