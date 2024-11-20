package repository

import (
	"context"
	"fmt"
	"task-tracker-server/internal/domain/user/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/minio/minio-go/v7"
)

func (r userRepo) RemoveAvatar(ctx context.Context, userID int, objectName string) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	err := r.Minio.Client.RemoveObject(ctx, r.Minio.Bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("user - repository - RemoveAvatar - r.Minio.Client.RemoveObject: %w", err)
	}

	sql, args, err := r.Builder.
		Update("users").
		SetMap(sq.Eq{
			"meta": &entity.UserMeta{Avatar: ""},
		}).
		Where(sq.Eq{"id": userID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("user - repository - RemoveAvatar - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("user - repository - RemoveAvatar - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
