package repository

import (
	"context"
	"projectly-server/internal/domain/team/entity"
	"projectly-server/pkg/postgres"
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
	GetStatisticData(ctx context.Context, teamID int) ([]entity.StatisticData, error)
	CheckUserInTeam(ctx context.Context, teamID, userID int) (bool, error)
	GetUserRole(ctx context.Context, teamID, userID int) (int, error)
	GetProjectTeamID(ctx context.Context, projectID int) (int, error)
	GetBoardTeamID(ctx context.Context, boardID int) (int, error)
	GetStatusTeamID(ctx context.Context, statusID int) (int, error)
	GetTaskTeamID(ctx context.Context, taskID int) (int, error)
}

const (
	_defaultConnTimeout = 5 * time.Second
)

type teamRepository struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) TeamRepository {
	return &teamRepository{pg}
}
