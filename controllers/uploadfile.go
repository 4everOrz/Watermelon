package controllers

import (
	"Watermelon/models"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BasicResponse{http.StatusBadRequest, "upload failed", nil})
		return
	}
	filename := header.Filename
	out, err := os.Create("files/" + filename) //存入项目目录中的files目录下
	if err != nil {
		log.Println("failed on opening the path")
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("failed on copying file")
	}
	c.JSON(http.StatusCreated, models.BasicResponse{http.StatusCreated, "upload successed", nil})
}
