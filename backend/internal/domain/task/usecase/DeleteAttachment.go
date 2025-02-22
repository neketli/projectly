package usecase

import (
	"context"
)

func (u *taskUseCase) DeleteAttachment(ctx context.Context, filename string) error {
	err := u.repo.DeleteAttachment(ctx, filename)
	if err != nil {
		u.logger.Error("task - usecase - DeleteAttachment - u.repo.DeleteAttachment: %s", err.Error())
		return err
	}

	return nil
}
