package routers

import "go-shopping/services"

func InitPublicRouter() {
	r := router.Group("/public")

	r.POST("/login", services.LoginHandler)
	r.POST("/register", services.RegisterHandler)
}
