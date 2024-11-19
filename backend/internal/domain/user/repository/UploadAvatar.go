package repository

import (
	"context"
	"fmt"
	"io"
	"task-tracker-server/internal/domain/user/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/minio/minio-go/v7"
)

func (r userRepo) UploadAvatar(ctx context.Context, user entity.User, reader io.Reader, filename string) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	objectName := fmt.Sprintf("avatars/%d/%s", user.ID, filename)

	_, err := r.Minio.Client.PutObject(ctx, r.Minio.Bucket, objectName, reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("user - repository - UploadAvatar - r.Minio.Client.PutObject: %w", err)
	}

	sql, args, err := r.Builder.
		Update("users").
		SetMap(sq.Eq{
			"meta": &entity.UserMeta{Avatar: objectName},
		}).
		Where(sq.Eq{"id": user.ID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("user - repository - UploadAvatar - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("user - repository - UploadAvatar - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
