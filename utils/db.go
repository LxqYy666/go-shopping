package utils

import (
	"fmt"
	"go-shopping/config"
	"go-shopping/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.Username, config.Password, config.DBname, config.Port)
	// fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Cart{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		panic("数据库初始化失败")
	}
	db.Create(&models.User{Username: "admin", Email: "example@xx.com", Password: "admin", Role: "admin", Status: "active"})

	DB = db
}
