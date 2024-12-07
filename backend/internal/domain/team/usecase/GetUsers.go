package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
)

func (u *teamUseCase) GetUsers(ctx context.Context, teamID int) ([]entity.TeamUser, error) {
	users, err := u.repo.GetUsers(ctx, teamID)
	if err != nil {
		u.logger.Error("team - usecase - GetUsers - u.repo.GetUsers: %s", err.Error())
		return nil, err
	}
	return users, nil
}
