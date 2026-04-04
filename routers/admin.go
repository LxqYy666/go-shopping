package routers

import "go-shopping/services"

func InitAdminRouter() {
	r := router.Group("/admin")

	categoryRouter := r.Group("/category")
	categoryRouter.GET("/list", services.CategoryListHandler)
	categoryRouter.POST("/add", services.CategoryAddHandler)

	productRouter := r.Group("/product")
	productRouter.GET("list", services.ProductListHandler)

}
