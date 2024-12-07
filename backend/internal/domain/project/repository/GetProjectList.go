package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/project/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r projectRepo) GetProjectList(ctx context.Context, teamID int) ([]entity.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "title", "description", "code", "team_id").
		From("project").
		Where(sq.Eq{"team_id": teamID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("project - repository - GetProjectList - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("project - repository - GetProjectList - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	projects, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Project])
	if err != nil {
		return nil, fmt.Errorf("project - repository - GetProjectList - pgx.CollectRows: %w", err)
	}

	return projects, nil
}
