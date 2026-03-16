package main

import (
	"fmt"
	"go-shopping/routers"
	"go-shopping/utils"
)

func main() {
	fmt.Println("这里是母婴商城项目")
	utils.InitDB()
	routers.InitRouter()
}
