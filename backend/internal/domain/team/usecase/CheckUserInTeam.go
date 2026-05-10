package usecase

import (
	"context"
)

// CheckUserInTeam checks if a user is a member of a team.
func (u *teamUseCase) CheckUserInTeam(ctx context.Context, teamID, userID int) (bool, error) {
	isUserInTeam, err := u.repo.CheckUserInTeam(ctx, teamID, userID)
	if err != nil {
		u.logger.Error("team - usecase - CheckUserInTeam - u.repo.CheckUserInTeam: %s", err.Error())
		return false, err
	}
	return isUserInTeam, nil
}
