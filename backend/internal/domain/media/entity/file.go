package entity

import (
	"time"

	"github.com/minio/minio-go/v7"
)

// FileInfo contains metadata about a file.
type FileInfo struct {
	Name         string
	Size         int64
	LastModified time.Time
	ContentType  string
	ETag         string
}

// File represents a file with its content reader.
type File struct {
	FileInfo FileInfo
	Reader   *minio.Object
}
