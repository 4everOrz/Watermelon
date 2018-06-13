package main

import (
	_ "Watermelon/common"
	"Watermelon/config"
	_ "Watermelon/db"
	_ "Watermelon/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.GetConf().DeBug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
