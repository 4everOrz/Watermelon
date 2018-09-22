package controllers

import (
	"Watermelon/common/jwt"
	_ "Watermelon/common/log"
	"Watermelon/models"
	"net/http"
	"security"
	"time"

	log "github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

//注册用户(字段空值校验由前端完成)
func (this *UserController) UserRegist(c *gin.Context) {
	var user models.User
	var notexist bool
	var err error
	if err = c.Bind(&user); err != nil {
		log.Error("Bind error:", err)
		return
	}
	user.Password = security.Md5(user.Password)
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateAt = user.CreatedAt
	notexist, _ = user.EntityExist()
	if notexist {
		if err := user.Create(); err != nil {
			log.Error("regist failed,error:", err)
			c.JSON(http.StatusBadRequest, BasicResponse{http.StatusBadRequest, "regist failed", nil})
		} else {
			c.JSON(http.StatusOK, BasicResponse{http.StatusOK, "regist success", nil})
		}
	} else {
		c.JSON(http.StatusBadRequest, BasicResponse{http.StatusBadRequest, "LoginName existed", nil})
	}
}

//用户登陆
func (this *UserController) UserLogin(c *gin.Context) {
	var user models.User
	var usermap = make(map[string]interface{})
	if err := c.Bind(&user); err != nil {
		log.Error("Bind error:", err)
	}
	userinfo, err := user.GetEntityByLoginName()
	if err != nil {
		log.Error("Error on GetEntityByLoginName :", err)
		return
	}
	if userinfo.Password == security.Md5(user.Password) {
		userinfo.Token = jwt.GenToken()
		userinfo.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		usermap["UserID"] = userinfo.UserID

		if err := userinfo.UpdateUser(); err != nil {
			log.Error("Error on UpdateUser :", err)
		}
		cookie := &http.Cookie{
			Name:     "token",
			Value:    userinfo.Token,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   10,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, BasicResponse{http.StatusOK, "login success", usermap})
	} else {
		c.JSON(http.StatusBadRequest, BasicResponse{http.StatusBadRequest, "login failed", nil})
	}
}
