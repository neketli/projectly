package usecase

import (
	"context"
	"task-tracker-server/internal/domain/team/entity"
	"task-tracker-server/internal/domain/team/repository"
	userEntity "task-tracker-server/internal/domain/user/entity"
	"task-tracker-server/pkg/logger"
)

type TeamUsecase interface {
	CreateTeam(ctx context.Context, team *entity.Team) error
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, teamID int) error
	GetTeamByUser(ctx context.Context, userID int) ([]entity.Team, error)
	GetUsers(ctx context.Context, teamID int) ([]userEntity.User, error)
	AddUserToTeam(ctx context.Context, teamID, userID int) error
	RemoveUserFromTeam(ctx context.Context, teamID, userID int) error
}

type teamUseCase struct {
	repo   repository.TeamRepository
	logger *logger.Logger
}

func New(r repository.TeamRepository, l *logger.Logger) TeamUsecase {
	return &teamUseCase{
		repo:   r,
		logger: l,
	}
}
