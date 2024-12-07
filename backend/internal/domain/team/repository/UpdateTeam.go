package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/team/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) UpdateTeam(ctx context.Context, team *entity.Team) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Update("team").
		SetMap(sq.Eq{
			"name":        team.Name,
			"description": team.Description,
		}).
		Where(sq.Eq{"id": team.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("team - repository - UpdateTeam - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("team - repository - UpdateTeam - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("team - repository - UpdateTeam - team not found")
	}

	return nil
}
