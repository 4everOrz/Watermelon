package controllers

import (
	"Watermelon/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	var usermap = make(map[string]interface{})
	var user models.User
	var err error
	user.Name = c.GetString("Name")
	user.Age = c.GetString("Age")
	usermap["UserID"], err = models.Addone(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BasicResponse{http.StatusBadRequest, "insertion failed", nil})
	} else {
		c.JSON(http.StatusOK, models.BasicResponse{http.StatusOK, "insertion success", usermap})
	}
}
