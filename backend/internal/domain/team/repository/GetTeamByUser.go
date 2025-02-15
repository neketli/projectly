package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/team/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r teamRepo) GetTeamByUser(ctx context.Context, userID int) ([]entity.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("team.id", "team.name", "team.description").
		From("team").
		Join("team_user tu on tu.team_id = team.id").
		Where(sq.Eq{"tu.user_id": userID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetTeamByUser - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetTeamByUser - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	teams, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Team])
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetTeamByUser - pgx.CollectRows: %w", err)
	}

	return teams, nil
}
