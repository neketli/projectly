package usecase

import (
	"context"
)

func (u *teamUseCase) GetStatusTeamID(ctx context.Context, statusID int) (int, error) {
	teamID, err := u.repo.GetStatusTeamID(ctx, statusID)
	if err != nil {
		u.logger.Error("team - usecase - GetStatusTeamID - u.repo.GetStatusTeamID: %s", err.Error())
		return 0, err
	}
	return teamID, nil
}
