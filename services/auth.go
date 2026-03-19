package services

import (
	"errors"
	"fmt"
	"go-shopping/dao"
	"go-shopping/net"
	"go-shopping/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context) {
	var loginBody net.LoginReq
	if err := c.BindJSON(&loginBody); err != nil {
		fmt.Println("绑定失败")
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
		return
	}

	if user, err := dao.HasAUser(loginBody); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "用户名或密码错误"))
		}
	} else {
		var exp int
		if loginBody.Remember {
			exp = 7 * 24 * 3600e9
		} else {
			exp = 24 * 3600e9
		}
		token, err := utils.GenerateJWT(jwt.SigningMethodHS256, exp, int(user.ID))
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
		} else {
			c.JSON(http.StatusOK, net.NewRes(http.StatusOK, net.LoginResData{Token: token}, "登陆成功"))
		}
	}
}

func RegisterHandler(c *gin.Context) {
	var registerBody net.RegisterReq
	if err := c.BindJSON(&registerBody); err != nil {
		c.JSON(http.StatusServiceUnavailable, net.NewRes(http.StatusServiceUnavailable, nil, "请求不可用"))
		return
	}

	if err := dao.CreateAUser(registerBody); err != nil {
		c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusBadRequest, nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, net.NewRes(http.StatusOK, nil, "注册成功"))
}
