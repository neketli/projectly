package usecase

import (
	"context"
	usersEntity "task-tracker-server/internal/domain/user/entity"
)

func (u *teamUseCase) GetUsers(ctx context.Context, teamID int) ([]usersEntity.User, error) {
	users, err := u.repo.GetUsers(ctx, teamID)
	if err != nil {
		u.logger.Error("team - usecase - GetUsers - u.repo.GetUsers: %s", err.Error())
		return nil, err
	}
	return users, nil
}
