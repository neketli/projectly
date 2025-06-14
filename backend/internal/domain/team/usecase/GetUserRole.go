package usecase

import (
	"context"
	"projectly-server/internal/domain/team/entity"
)

func (u *teamUseCase) GetUserRole(ctx context.Context, teamID, userID int) (*entity.Role, error) {
	roleID, err := u.repo.GetUserRole(ctx, teamID, userID)
	if err != nil {
		u.logger.Error("team - usecase - GetUserRole - u.repo.GetUserRole: %s", err.Error())
		return nil, err
	}

	switch roleID {
	case entity.RoleOwner.ID:
		return entity.RoleOwner, nil
	case entity.RoleEditor.ID:
		return entity.RoleEditor, nil
	case entity.RoleDeveloper.ID:
		return entity.RoleDeveloper, nil
	case entity.RoleUser.ID:
		return entity.RoleUser, nil
	default:
		return nil, nil
	}
}
