package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitPublicRouter() {
	r := router.Group("/public")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "这是public路由",
		})
	})
}
