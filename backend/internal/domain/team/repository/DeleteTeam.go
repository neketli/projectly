package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) DeleteTeam(ctx context.Context, teamID int) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Delete("teams").
		Where(sq.Eq{"id": teamID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("team - repository - DeleteTeam - r.Builder: %w", err)
	}

	result, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("team - repository - DeleteTeam - r.Pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("team - repository - DeleteTeam - team not found")
	}

	return nil
}
