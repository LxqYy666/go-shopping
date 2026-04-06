package services

import (
	"fmt"
	"go-shopping/dao"
	"go-shopping/net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CartListHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	cartItems, err := dao.GetCartItems(userID.(uint))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "获取购物车失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, cartItems, "获取购物车成功"))
}

func CartAddHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	var req net.AddToCartReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "请求参数错误"))
		return
	}

	if err := dao.AddToCart(userID.(uint), req); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "添加到购物车成功"))
}

func CartUpdateHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	id := c.Param("id")
	var cartID uint
	if _, err := fmt.Sscanf(id, "%d", &cartID); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "无效的购物车ID"))
		return
	}

	var req net.UpdateCartReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "请求参数错误"))
		return
	}

	if err := dao.UpdateCartItem(userID.(uint), cartID, req); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "更新购物车成功"))
}

func CartDeleteHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	id := c.Param("id")
	var cartID uint
	if _, err := fmt.Sscanf(id, "%d", &cartID); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "无效的购物车ID"))
		return
	}

	if err := dao.DeleteCartItem(userID.(uint), cartID); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "删除成功"))
}

func CreateOrderHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	var req net.CreateOrderReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, net.NewRes(http.StatusBadRequest, nil, "请求参数错误"))
		return
	}

	if err := dao.CreateOrder(userID.(uint), req); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "订单创建成功"))
}
