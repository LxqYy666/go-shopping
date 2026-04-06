package services

import (
	"fmt"
	"go-shopping/dao"
	"go-shopping/net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryAddHandler(c *gin.Context) {
	var addCategoryReq net.AddCategoryReq
	if err := c.BindJSON(&addCategoryReq); err != nil {
		fmt.Println("绑定失败")
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
		return
	}

	if err := dao.AddCategory(addCategoryReq); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "添加成功"))
}

func CategoryListHandler(c *gin.Context) {

	if categoryList, err := dao.GetCategoryInfo(); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
	} else {
		c.JSON(http.StatusOK, net.NewRes(http.StatusOK, categoryList, "获取种类信息成功"))
	}

}

func ProductListHandler(c *gin.Context) {
	if productList, err := dao.GetProductList(); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
	} else {
		c.JSON(http.StatusOK, net.NewRes(http.StatusOK, productList, "获取商品信息成功"))
	}
}

func UserListHandler(c *gin.Context) {
	if userList, err := dao.GetUserList(); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
	} else {
		c.JSON(http.StatusOK, net.NewRes(http.StatusOK, userList, "获取用户信息成功"))
	}
}

func ProductAddHandler(c *gin.Context) {
	var req net.AddProductReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "请求参数错误"))
		return
	}

	if err := dao.AddProduct(req); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "添加商品失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "添加商品成功"))
}

func ProductUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var productID uint
	if _, err := fmt.Sscanf(id, "%d", &productID); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "无效的商品ID"))
		return
	}

	var req net.UpdateProductReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "请求参数错误"))
		return
	}

	if err := dao.UpdateProduct(productID, req); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "更新商品失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "更新商品成功"))
}

func ProductDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	var productID uint
	if _, err := fmt.Sscanf(id, "%d", &productID); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "无效的商品ID"))
		return
	}

	if err := dao.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "删除商品失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "删除商品成功"))
}

func OrderListHandler(c *gin.Context) {
	orderList, err := dao.GetOrderList()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "获取订单列表失败"))
		return
	}
	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, orderList, "获取订单列表成功"))
}

func OrderUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var orderID uint
	if _, err := fmt.Sscanf(id, "%d", &orderID); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "无效的订单ID"))
		return
	}

	var req net.UpdateOrderReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "请求参数错误"))
		return
	}

	if err := dao.UpdateOrder(orderID, req); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "更新订单失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "更新订单成功"))
}
