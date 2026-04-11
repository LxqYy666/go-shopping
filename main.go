package main

import (
	"fmt"
	"go-shopping/routers"
	"go-shopping/services"
	"go-shopping/utils"
)

func main() {
	fmt.Println("这里是母婴商城项目")
	utils.InitDB()
	if err := utils.InitRedis(); err != nil {
		fmt.Printf("Redis连接失败：%v\n", err)
	} else {
		fmt.Println("Redis连接成功")
	}
	services.StartOrderTimeoutWorker()
	routers.InitRouter()
}
