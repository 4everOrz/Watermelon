package controllers

import (
	"Watermelon/common/log"
	"Watermelon/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	var user models.User
	var err error
	user.Name = c.PostForm("Name")
	user.Age, _ = strconv.Atoi(c.PostForm("Age"))
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = user.Create()
	if err != nil {
		log.Logger.Error("insertion failed,error:", err)
		c.JSON(http.StatusBadRequest, models.BasicResponse{http.StatusBadRequest, "insertion failed", nil})
	} else {
		c.JSON(http.StatusOK, models.BasicResponse{http.StatusOK, "insertion success", nil})
	}
}
