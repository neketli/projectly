package usecase

import (
	"context"
	"projectly-server/internal/domain/media/entity"
)

func (u *mediaUseCase) GetFile(ctx context.Context, filename string) (*entity.File, error) {
	file, err := u.repo.GetFile(ctx, filename)
	if err != nil {
		u.logger.Error("media - usecase - GetFile - u.repo.GetFile: %s", err.Error())
		return nil, err
	}

	return file, nil
}
