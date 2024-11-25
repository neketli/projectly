package repository

import (
	"context"
	"errors"
	"fmt"

	"task-tracker-server/internal/domain/project/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r projectRepo) GetProject(ctx context.Context, projectID int) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "title", "description", "code", "team_id").
		From("project").
		Where(sq.Eq{"id": projectID}).
		ToSql()
	if err != nil {
		return entity.Project{}, fmt.Errorf("project - repository - GetProject - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.Project{}, fmt.Errorf("project - repository - GetProject - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	projects, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Project])
	if err != nil {
		return entity.Project{}, fmt.Errorf("project - repository - GetProject - pgx.CollectRows: %w", err)
	}

	if len(projects) == 0 {
		return entity.Project{}, errors.New("no project found")
	}

	return projects[0], nil
}
