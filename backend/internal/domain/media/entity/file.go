package entity

import (
	"time"

	"github.com/minio/minio-go/v7"
)

type FileInfo struct {
	Name         string
	Size         int64
	LastModified time.Time
	ContentType  string
	ETag         string
}

type File struct {
	FileInfo FileInfo
	Reader   *minio.Object
}
