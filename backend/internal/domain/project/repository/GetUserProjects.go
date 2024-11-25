package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/project/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r projectRepo) GetUserProjects(ctx context.Context, userID int) ([]entity.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("p.id", "p.title", "p.description", "p.code", "p.team_id").
		From("project p").
		Join("team_user tu on tu.team_id = p.team_id").
		Where(sq.Eq{"tu.user_id": userID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("project - repository - GetUserProjects - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("project - repository - GetUserProjects - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	projects, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Project])
	if err != nil {
		return nil, fmt.Errorf("project - repository - GetUserProjects - pgx.CollectRows: %w", err)
	}

	return projects, nil
}
