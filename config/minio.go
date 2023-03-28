package config

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMiniO(cgf Config) (*minio.Client, error) {
	minioClient, err := minio.New(cgf.StorageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cgf.StorageUser, cgf.StoragePassword, ""),
		Secure: cgf.StorageSSL,
	})
	return minioClient, err
}

// Create a bucket at region 'us-east-1' with object locking enabled.
func CreateMakeBucket(ctx context.Context, cgf Config, minioClient *minio.Client) (err error) {
	return minioClient.MakeBucket(ctx, cgf.StorageBucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
}
