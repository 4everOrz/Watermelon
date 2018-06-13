package routers

import (
	"Watermelon/config"
	"Watermelon/controllers"
	"Watermelon/webscoket"

	"github.com/gin-gonic/gin"
)

func init() {

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery()) //全局中间件
	router_api(router).Run(":" + config.GetConf().HandlePort)
}

//api
func router_api(router *gin.Engine) *gin.Engine {
	router1 := router.Group("/v1")                  //分组路由
	router1.POST("/form_post", controllers.UserAdd) //访问示例: 地址+/v1/form_post?
	router1.POST("/upload", controllers.UploadFile)
	router1.GET("/ws", webscoket.WebscoketJoin)
	return router
}
