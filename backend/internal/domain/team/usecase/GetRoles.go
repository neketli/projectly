package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
)

func (u *teamUseCase) GetRoles(ctx context.Context) ([]entity.Role, error) {
	roles, err := u.repo.GetRoles(ctx)
	if err != nil {
		u.logger.Error("team - usecase - GetRoles - u.repo.GetRoles: %s", err.Error())
		return nil, err
	}
	return roles, nil
}
