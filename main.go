package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"

	"github.com/mrkonxyz/config"
)

func main() {

	cgf := config.LoadConFig(".")
	//ctx := context.Background()
	minioClient, err := config.ConnectMiniO(cgf)
	if err != nil {
		log.Fatalln(err)
	}

	// err = config.CreateMakeBucket(ctx, cgf, minioClient)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	r := gin.Default()

	r.POST("/upload", func(ctx *gin.Context) {
		userMetaData := map[string]string{"x-amz-acl": "public-read"}
		img, _ := ctx.FormFile("img")
		file, _ := img.Open()
		contentType := img.Header.Get("Content-Type")
		minioResult, err := minioClient.PutObject(ctx.Request.Context(),
			cgf.StorageBucketName,
			img.Filename,
			file, -1,
			minio.PutObjectOptions{
				ContentType:  contentType,
				UserMetadata: userMetaData,
			},
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, minioResult.Location)

	})

	r.Run()

}
