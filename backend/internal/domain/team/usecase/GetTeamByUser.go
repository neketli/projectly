package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
)

func (u *teamUseCase) GetTeamByUser(ctx context.Context, userID int) ([]entity.Team, error) {
	teams, err := u.repo.GetTeamByUser(ctx, userID)
	if err != nil {
		u.logger.Error("team - usecase - GetTeamByUser - u.repo.GetTeamByUser: %s", err.Error())
		return nil, err
	}
	return teams, nil
}
