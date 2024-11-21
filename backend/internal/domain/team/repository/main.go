package repository

import (
	"context"
	"task-tracker-server/internal/domain/team/entity"
	userEntity "task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/pkg/postgres"
	"time"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *entity.Team) error
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, teamID int) error
	GetTeamByUser(ctx context.Context, userID int) ([]entity.Team, error)
	GetUsers(ctx context.Context, teamID int) ([]userEntity.User, error)
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
