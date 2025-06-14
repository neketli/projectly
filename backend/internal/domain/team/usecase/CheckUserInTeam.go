package usecase

import (
	"context"
)

func (u *teamUseCase) CheckUserInTeam(ctx context.Context, teamID, userID int) (bool, error) {
	isUserInTeam, err := u.repo.CheckUserInTeam(ctx, teamID, userID)
	if err != nil {
		u.logger.Error("team - usecase - CheckUserInTeam - u.repo.CheckUserInTeam: %s", err.Error())
		return false, err
	}
	return isUserInTeam, nil
}
