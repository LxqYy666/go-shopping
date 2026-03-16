package routers

import (
	"fmt"
	"go-shopping/config"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter() {
	router = gin.Default()
	InitPublicRouter()
	InitUserRouter()
	InitAdminRouter()
	router.Run(fmt.Sprintf("%s:%s", config.ServerIp, config.ServerPort))
}
