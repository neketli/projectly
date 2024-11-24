package repository

import (
	"context"
	"task-tracker-server/internal/domain/team/entity"
	"task-tracker-server/pkg/postgres"
	"time"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *entity.Team) error
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, teamID int) error
	GetTeam(ctx context.Context, teamID int) (entity.Team, error)
	GetRoles(ctx context.Context) ([]entity.Role, error)
	SetRole(ctx context.Context, teamID, userID, roleID int) error
	GetTeamByUser(ctx context.Context, userID int) ([]entity.Team, error)
	GetUsers(ctx context.Context, teamID int) ([]entity.TeamUser, error)
	AddUserToTeam(ctx context.Context, teamID, userID int) error
	RemoveUserFromTeam(ctx context.Context, teamID, userID int) error
}

const (
	_defaultConnTimeout = 5 * time.Second
)

type teamRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) TeamRepository {
	return teamRepo{pg}
}
