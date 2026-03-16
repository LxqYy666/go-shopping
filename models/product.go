package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID uint    `json:"category_id"`
	Name       string  `gorm:"not null" json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `gorm:"not null" json:"price"`
	Stock      int     `gorm:"default:0" json:"stock"`
	ImageUrl   string  `json:"image_url"`
	SoldCount  int     `gorm:"default:0" json:"sold_count"`
	Status     string  `gorm:"default:'on';check:status in ('on','off')" json:"status"`
}
