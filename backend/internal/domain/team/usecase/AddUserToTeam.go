package usecase

import (
	"context"
)

func (u *teamUseCase) AddUserToTeam(ctx context.Context, teamID, userID int) error {
	err := u.repo.AddUserToTeam(ctx, teamID, userID)
	if err != nil {
		u.logger.Error("team - usecase - CreateTeam - u.repo.AddUserToTeam: %s", err.Error())
		return err
	}
	return nil
}
