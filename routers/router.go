package routers

import (
	"Watermelon/common/config"

	"Watermelon/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	//模式选择
	if config.GetString("SetMode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//路由
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery()) //全局中间件
	router_api(router).Run(":" + config.GetString("ListenPort"))
}

//api
func router_api(router *gin.Engine) *gin.Engine {
	usercontroller := &controllers.UserController{}
	basecontroller := &controllers.BaseController{}
	productcontroller := &controllers.ProductController{}
	router1 := router.Group("/v1")                           //分组路由
	router1.POST("/regist", usercontroller.UserRegist)       //json
	router1.POST("/login", usercontroller.UserLogin)         //json
	router1.POST("/upload", basecontroller.UploadFile)       //formdata
	router1.POST("/addproduct", productcontroller.InsertOne) //formdata
	return router
}
