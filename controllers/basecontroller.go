package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//response结构体
type BasicResponse struct {
	Code int
	Msg  string
	Data interface{}
}
type BaseController struct {
}

//文件上传
func (this *BaseController) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusBadRequest, BasicResponse{http.StatusBadRequest, "upload failed", nil})
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
	c.JSON(http.StatusCreated, BasicResponse{http.StatusCreated, "upload successed", nil})
}
