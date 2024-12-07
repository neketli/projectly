package repository

import (
	"context"
	"errors"
	"fmt"

	"projectly-server/internal/domain/team/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r teamRepo) GetTeam(ctx context.Context, teamID int) (entity.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "name", "description").
		From("team").
		Where(sq.Eq{"id": teamID}).
		ToSql()
	if err != nil {
		return entity.Team{}, fmt.Errorf("team - repository - GetTeam - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.Team{}, fmt.Errorf("team - repository - GetTeam - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	teams, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Team])
	if err != nil {
		return entity.Team{}, fmt.Errorf("team - repository - GetTeam - pgx.CollectRows: %w", err)
	}

	if len(teams) == 0 {
		return entity.Team{}, errors.New("no team found")
	}

	return teams[0], nil
}
