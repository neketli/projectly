package usecase

import (
	"context"
	"task-tracker-server/internal/domain/status/entity"
)

func (u *statusUseCase) UpdateStatus(ctx context.Context, status *entity.Status) error {
	err := u.repo.UpdateStatus(ctx, status)
	if err != nil {
		u.logger.Error("status - usecase - UpdateStatus - u.repo.UpdateStatus: %s", err.Error())
		return err
	}
	return nil
}
