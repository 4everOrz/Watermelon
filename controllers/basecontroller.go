package controllers

import (
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/alecthomas/log4go"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	key []byte = []byte("-jwt-gree@gree.com")
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
		log.Error("failed on opening the path")
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Error("failed on copying file")
	}
	c.JSON(http.StatusCreated, BasicResponse{http.StatusCreated, "upload successed", nil})
}

//jwt
// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "gree132121",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Error("create token err:", err)
		return ""
	}
	return ss
}

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Error("parase with claims failed.", err)
		return false
	}
	return true
}
func FailOnErr(err error, msg string) {
	if err != nil {
		log.Error("%s:%s", msg, err)
	}
}
