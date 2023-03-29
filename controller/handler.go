package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/mrkonxyz/config"
)

type HandlerFunc struct {
	cgf *config.Config
}

func NewHandlerFunc(cgf *config.Config) HandlerFunc {
	return HandlerFunc{
		cgf: cgf,
	}
}

func (h *HandlerFunc) UploadPicture(minioClient *minio.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		img, _ := ctx.FormFile("img")
		file, _ := img.Open()
		contentType := img.Header.Get("Content-Type")

		result, err := UploadPicture(minioClient, ctx, contentType, file, h.cgf.StorageBucketName, img.Filename)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, result.Location)
	}
}
