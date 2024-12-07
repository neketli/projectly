package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
	"projectly-server/internal/domain/team/repository"
	"projectly-server/pkg/logger"
)

type TeamUseCase interface {
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

type teamUseCase struct {
	repo   repository.TeamRepository
	logger *logger.Logger
}

func New(r repository.TeamRepository, l *logger.Logger) TeamUseCase {
	return &teamUseCase{
		repo:   r,
		logger: l,
	}
}
