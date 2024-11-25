package repository

import (
	"context"
	"errors"
	"fmt"

	"task-tracker-server/internal/domain/project/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r projectRepo) GetProjectByCode(ctx context.Context, teamID int, code string) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "title", "description", "code", "team_id").
		From("project").
		Where(sq.And{
			sq.Eq{"team_id": teamID},
			sq.Eq{"code": code},
		}).
		ToSql()
	if err != nil {
		return entity.Project{}, fmt.Errorf("project - repository - GetProjectByCode - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.Project{}, fmt.Errorf("project - repository - GetProjectByCode - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	projects, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Project])
	if err != nil {
		return entity.Project{}, fmt.Errorf("project - repository - GetProjectByCode - pgx.CollectRows: %w", err)
	}

	if len(projects) == 0 {
		return entity.Project{}, errors.New("no project found")
	}

	return projects[0], nil
}
