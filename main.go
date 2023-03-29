package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mrkonxyz/config"
	"github.com/mrkonxyz/controller"
)

func main() {

	cgf := config.LoadConFig(".")
	ctx := context.Background()
	minioClient, err := config.ConnectMiniO(cgf, ctx)

	if err != nil {
		log.Fatalln(err)
	}
	h := controller.NewHandlerFunc(&cgf)

	r := gin.Default()
	r.POST("/upload", h.UploadPicture(minioClient))

	r.Run()

}
