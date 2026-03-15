package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `gorm:"not null" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}
