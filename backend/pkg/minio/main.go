package minio

import (
	"fmt"
	"projectly-server/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Client *minio.Client
	Bucket string
}

func New(config config.S3) (*Minio, error) {
	minioClient, err := minio.New(
		config.Host,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				config.AccessKey,
				config.SecretKey,
				"",
			),
			Secure: config.Secure,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("minio - New - minio.New: %w", err)

	}

	return &Minio{
		Client: minioClient,
		Bucket: config.Bucket,
	}, nil
}
