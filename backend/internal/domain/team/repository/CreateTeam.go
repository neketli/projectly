package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/team/entity"
)

func (r teamRepo) CreateTeam(ctx context.Context, team *entity.Team) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("team").
		Columns(
			"name",
			"description",
		).
		Values(
			team.Name,
			team.Description,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("team - repository - CreateTeam - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("team - repository - CreateTeam - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return fmt.Errorf("team - repository - CreateTeam - rows.Scan: %w", err)
		}
		team.ID = id
	}

	return nil
}
