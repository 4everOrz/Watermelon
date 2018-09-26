package controllers

import (
	"Watermelon/models"
	"net/http"
	"strconv"
	"time"

	log "github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	BaseController
}

func (this *ProductController) InsertOne(c *gin.Context) {
	var product models.Product
	product.CreatorID, _ = strconv.ParseInt(c.PostForm("UserID"), 10, 64)
	token, err := c.Cookie("token")
	if err != nil {
		log.Error("token error")
		c.JSON(http.StatusUnauthorized, BasicResponse{http.StatusUnauthorized, "auth failed", nil})
		return
	}
	if AuthToken(product.CreatorID, token) {
		product.ProductName = c.PostForm("ProductName")
		product.ProductNumber = c.PostForm("ProductNumber")
		product.Price = c.PostForm("Price")
		product.CreateAt = time.Now().Format("2006-01-02 15:04:05")
		product.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
		err = product.InsertOne()
		FailOnErr(err, "insert new record error!")
		c.JSON(http.StatusOK, BasicResponse{http.StatusOK, "add success", nil})
	} else {
		c.JSON(http.StatusBadRequest, BasicResponse{http.StatusBadRequest, "add failed", nil})
	}
}
