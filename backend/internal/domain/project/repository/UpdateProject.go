package repository

import (
	"context"
	"fmt"

	"projectly-server/internal/domain/project/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r projectRepo) UpdateProject(ctx context.Context, project *entity.Project) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Update("project").
		SetMap(sq.Eq{
			"title":       project.Title,
			"description": project.Description,
		}).
		Where(sq.Eq{"id": project.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("project - repository - UpdateProject - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("project - repository - UpdateProject - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("project - repository - UpdateProject - project not found")
	}

	return nil
}
