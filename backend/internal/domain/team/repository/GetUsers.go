package repository

import (
	"context"
	"fmt"

	"task-tracker-server/internal/domain/team/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r teamRepo) GetUsers(ctx context.Context, teamID int) ([]entity.TeamUser, error) {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	sql, args, err := r.Builder.
		Select(
			"tu.user_id",
			"u.email",
			"u.name",
			"u.surname",
			"r.id",
			"r.role_name",
		).
		From("team_user tu").
		Join("users u on u.id = tu.user_id").
		LeftJoin("team_roles tr on tr.team_id = tu.team_id AND tr.user_id = tu.user_id").
		LeftJoin("roles r on r.id = tr.role_id").
		Where(sq.Eq{"tu.team_id": teamID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetUsers - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("team - repository - GetUsers - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var users []entity.TeamUser
	for rows.Next() {
		var user entity.TeamUser
		var role entity.Role
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Surname,
			&role.ID,
			&role.RoleName,
		); err != nil {
			return nil, fmt.Errorf("team - repository - GetUsers - rows.Scan: %w", err)
		}
		user.Role = role
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("team - repository - GetUsers - no users found")
	}

	return users, nil
}
