package usecase

import (
	"context"
	"projectly-server/internal/domain/status/entity"
)

func (u *statusUseCase) CreateStatus(ctx context.Context, status *entity.Status) error {
	err := u.repo.CreateStatus(ctx, status)
	if err != nil {
		u.logger.Error("status - usecase - CreateStatus - u.repo.CreateStatus: %s", err.Error())
		return err
	}
	return nil
}
