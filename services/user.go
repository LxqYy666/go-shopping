package services

import (
	"go-shopping/dao"
	"go-shopping/net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserOrdersHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	orders, err := dao.GetUserOrders(userID.(uint))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "获取订单失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, orders, "获取订单成功"))
}

func UserInfoHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未认证"))
		return
	}

	user, err := dao.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "获取用户信息失败"))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, user, "获取用户信息成功"))
}
