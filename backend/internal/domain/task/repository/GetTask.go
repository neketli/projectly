package repository

import (
	"context"
	"errors"
	"fmt"

	"task-tracker-server/internal/domain/task/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r taskRepo) GetTask(ctx context.Context, taskID int) (entity.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select("id", "project_id", "title").
		From("task").
		Where(sq.Eq{"id": taskID}).
		ToSql()
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - GetTask - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - GetTask - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Task])
	if err != nil {
		return entity.Task{}, fmt.Errorf("task - repository - GetTask - pgx.CollectRows: %w", err)
	}

	if len(tasks) == 0 {
		return entity.Task{}, errors.New("no task found")
	}

	return tasks[0], nil
}
