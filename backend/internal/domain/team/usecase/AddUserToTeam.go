package usecase

import (
	"context"
)

// AddUserToTeam adds a user to a team.
func (u *teamUseCase) AddUserToTeam(ctx context.Context, teamID, userID int) error {
	err := u.repo.AddUserToTeam(ctx, teamID, userID)
	if err != nil {
		u.logger.Error("team - usecase - AddUserToTeam - u.repo.AddUserToTeam: %s", err.Error())
		return err
	}
	return nil
}
