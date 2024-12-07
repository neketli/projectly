package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
)

func (u *teamUseCase) GetTeam(ctx context.Context, teamID int) (entity.Team, error) {
	team, err := u.repo.GetTeam(ctx, teamID)
	if err != nil {
		u.logger.Error("team - usecase - GetTeam - u.repo.GetTeam: %s", err.Error())
		return entity.Team{}, err
	}
	return team, nil
}
