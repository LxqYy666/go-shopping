package routers

import (
	"go-shopping/middlewares"
	"go-shopping/services"
)

func InitUserRouter() {
	r := router.Group("/user", middlewares.AuthMiddleware())

	// 购物车
	r.GET("/cart", services.CartListHandler)
	r.POST("/cart", services.CartAddHandler)
	r.PUT("/cart/:id", services.CartUpdateHandler)
	r.DELETE("/cart/:id", services.CartDeleteHandler)

	// 订单
	r.POST("/order", services.CreateOrderHandler)
	r.GET("/orders", services.UserOrdersHandler)

	// 用户信息
	r.GET("/info", services.UserInfoHandler)
}
