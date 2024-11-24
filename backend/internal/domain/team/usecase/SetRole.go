package usecase

import (
	"context"
)

func (u *teamUseCase) SetRole(ctx context.Context, teamID, userID, roleID int) error {
	err := u.repo.SetRole(ctx, teamID, userID, roleID)
	if err != nil {
		u.logger.Error("team - usecase - SetRole - u.repo.SetRole: %s", err.Error())
		return err
	}
	return nil
}
