package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint     `json:"user_id"`
	ProductID uint     `json:"product_id"`
	Quantity  int      `gorm:"not null;check:quantity > 0" json:"quantity"`
	User      *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Product   *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
