package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/minio/minio-go/v7"
)

func (r taskRepo) DeleteAttachment(ctx context.Context, objectName string) error {
	ctx, cancel := context.WithTimeout(ctx, _defaultConnTimeout)
	defer cancel()

	err := r.Minio.Client.RemoveObject(ctx, r.Minio.Bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("user - repository - DeleteAttachment - r.Minio.Client.RemoveObject: %w", err)
	}

	sql, args, err := r.Builder.
		Delete("attachment").
		Where(sq.Eq{"name": objectName}).
		ToSql()

	if err != nil {
		return fmt.Errorf("user - repository - DeleteAttachment - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("user - repository - DeleteAttachment - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	return nil
}
