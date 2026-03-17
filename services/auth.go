package services

import (
	"errors"
	"go-shopping/dao"
	"go-shopping/net"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context) {
	var loginBody net.LoginReq
	if err := c.BindJSON(&loginBody); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
	}

	if user, err := dao.HasAUser(loginBody); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "用户名或密码错误"))
		}
	} else {
		c.JSON(http.StatusOK, net.NewRes(http.StatusOK, user.ID, "登陆成功"))
	}
}
