package routers

import (
	"go-shopping/middlewares"
	"go-shopping/services"
)

func InitAdminRouter() {
	r := router.Group("/admin", middlewares.AuthMiddleware())

	categoryRouter := r.Group("/category")
	categoryRouter.GET("/list", services.CategoryListHandler)
	categoryRouter.POST("/add", services.CategoryAddHandler)

	productRouter := r.Group("/product")
	productRouter.GET("/list", services.ProductListHandler)
	productRouter.POST("/add", services.ProductAddHandler)
	productRouter.PUT("/:id", services.ProductUpdateHandler)
	productRouter.DELETE("/:id", services.ProductDeleteHandler)

	userRouter := r.Group("/user")
	userRouter.GET("/list", services.UserListHandler)

	orderRouter := r.Group("/order")
	orderRouter.GET("/list", services.OrderListHandler)
	orderRouter.PUT("/:id", services.OrderUpdateHandler)
}
