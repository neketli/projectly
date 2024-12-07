package usecase

import (
	"context"
	"projectly-server/internal/domain/status/entity"
)

func (u *statusUseCase) UpdateStatus(ctx context.Context, status *entity.Status, oldOrder *int) error {
	if oldOrder != nil && *oldOrder != status.Order {
		err := u.repo.UpdateOrders(ctx, status.BoardID, *oldOrder, status.Order)
		if err != nil {
			u.logger.Error("status - usecase - UpdateStatus - u.repo.UpdateOrders: %s", err.Error())
			return err
		}
	}
	err := u.repo.UpdateStatus(ctx, status)
	if err != nil {
		u.logger.Error("status - usecase - UpdateStatus - u.repo.UpdateStatus: %s", err.Error())
		return err
	}
	return nil
}
