package usecase

import (
	"context"
)

// RemoveUserFromTeam removes a user from a team.
func (u *teamUseCase) RemoveUserFromTeam(ctx context.Context, teamID, userID int) error {
	err := u.repo.RemoveUserFromTeam(ctx, teamID, userID)
	if err != nil {
		u.logger.Error("team - usecase - CreateTeam - u.repo.RemoveUserFromTeam: %s", err.Error())
		return err
	}
	return nil
}
