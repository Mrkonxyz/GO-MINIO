package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mrkonxyz/config"
)

func main() {

	cgf := config.LoadConFig(".")

	// Initialize minio client object.
	minioClient, err := minio.New(cgf.StorageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cgf.StorageUser, cgf.StoragePassword, ""),
		Secure: cgf.StorageSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// Create a bucket at region 'us-east-1' with object locking enabled.
	err = minioClient.MakeBucket(context.Background(), cgf.StorageBucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("Successfully created mybucket.")

}
