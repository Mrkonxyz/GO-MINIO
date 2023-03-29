package controller

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func UploadPicture(minioClient *minio.Client, ctx *gin.Context, contentType string, file multipart.File, BucketName string, Filename string) (minioResult minio.UploadInfo, err error) {
	userMetaData := map[string]string{"x-amz-acl": "public-read"}

	return minioClient.PutObject(ctx.Request.Context(),
		BucketName,
		Filename,
		file, -1,
		minio.PutObjectOptions{
			ContentType:  contentType,
			UserMetadata: userMetaData,
		},
	)
}
