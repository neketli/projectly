package minio

import (
	"fmt"
	"projectly-server/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Minio provides S3-compatible object storage client.
type Minio struct {
	Client *minio.Client
	Bucket string
}

// New creates a new Minio client.
func New(cfg config.S3) (*Minio, error) {
	minioClient, err := minio.New(
		cfg.Host,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				cfg.AccessKey,
				cfg.SecretKey,
				"",
			),
			Secure: cfg.Secure,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("minio - New - minio.New: %w", err)

	}

	return &Minio{
		Client: minioClient,
		Bucket: cfg.Bucket,
	}, nil
}
