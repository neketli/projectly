package repository

import (
	"context"
	"fmt"
	"projectly-server/internal/domain/media/entity"

	"github.com/minio/minio-go/v7"
)

func (r *mediaRepo) GetFile(ctx context.Context, filename string) (*entity.File, error) {
	objInfo, err := r.Minio.Client.StatObject(ctx, r.Bucket, filename, minio.StatObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("media - repository - GetFile - r.Minio.Client.StatObject: %w", err)
	}

	object, err := r.Minio.Client.GetObject(ctx, r.Minio.Bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("media - repository - GetFile - r.Minio.Client.GetObject: %w", err)
	}

	fileContent := &entity.File{
		FileInfo: entity.FileInfo{
			Name:         filename,
			Size:         objInfo.Size,
			LastModified: objInfo.LastModified,
			ContentType:  objInfo.ContentType,
			ETag:         objInfo.ETag,
		},
		Reader: object,
	}

	return fileContent, nil
}
