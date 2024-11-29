package usecase

import (
	"context"
	"task-tracker-server/internal/domain/status/entity"
)

func (u *statusUseCase) GetStatusList(ctx context.Context, boardID int) ([]entity.Status, error) {
	statuses, err := u.repo.GetStatusList(ctx, boardID)
	if err != nil {
		u.logger.Error("status - usecase - GetStatusList - u.repo.GetStatusList: %s", err.Error())
		return nil, err
	}
	return statuses, nil
}
