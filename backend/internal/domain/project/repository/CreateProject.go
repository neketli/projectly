package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/project/entity"
)

func (r projectRepo) CreateProject(ctx context.Context, project *entity.Project) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Insert("project").
		Columns(
			"title",
			"description",
			"code",
			"team_id",
		).
		Values(
			project.Title,
			project.Description,
			project.Code,
			project.TeamID,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("project - repository - CreateProject - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("project - repository - CreateProject - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("project - repository - CreateProject - project not found")
	}

	var id int
	if err = rows.Scan(&id); err != nil {
		return fmt.Errorf("project - repository - CreateProject - rows.Scan: %w", err)
	}

	project.ID = id
	return nil
}
