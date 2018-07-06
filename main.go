package main

import (
	"Watermelon/common/config"
	_ "Watermelon/common/log"
	_ "Watermelon/db/mysql"
	_ "Watermelon/db/redis"
	_ "Watermelon/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.GetString("SetMode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
