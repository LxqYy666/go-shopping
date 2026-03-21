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
