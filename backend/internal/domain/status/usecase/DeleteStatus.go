package usecase

import (
	"context"
)

func (u *statusUseCase) DeleteStatus(ctx context.Context, statusID, order int) error {
	err := u.repo.DeleteStatus(ctx, statusID, order)
	if err != nil {
		u.logger.Error("status - usecase - DeleteStatus - u.repo.DeleteStatus: %s", err.Error())
		return err
	}
	return nil
}
