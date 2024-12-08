package repository

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/team/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) GetStatisticData(ctx context.Context, teamID int) ([]entity.StatisticData, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select(
			"p.id",
			"p.code",
			"COUNT(t.id) AS total_tasks_count",
			"COUNT(t.finished_at) AS completed_tasks_count",
			"AVG(EXTRACT(EPOCH FROM (t.finished_at - t.created_at))/3600) AS avg_task_duration").
		From("project p").
		Join("board b ON p.id = b.project_id").
		Join("status s ON b.id = s.board_id").
		Join("task t ON s.id = t.status_id").
		Where(sq.Eq{"p.team_id": teamID}).
		GroupBy("p.id", "p.code").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetStatisticData - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetStatisticData - r.Pool.Exec: %w", err)
	}
	defer rows.Close()

	data := make([]entity.StatisticData, 0)
	for rows.Next() {
		var row entity.StatisticData
		if err := rows.Scan(
			&row.ID,
			&row.Code,
			&row.TotalTasksCount,
			&row.CompletedTasksCount,
			&row.AvgTaskDuration,
		); err != nil {
			return nil, fmt.Errorf("team - repository - GetStatisticData - rows.Scan: %w", err)
		}
		data = append(data, row)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("team - repository - GetStatisticData - rows.Err: %w", err)
	}

	return data, nil
}
