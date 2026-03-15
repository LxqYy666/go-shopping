package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"unique;not null" json:"username"`
	Email    string  `gorm:"unique;not null" json:"email"`
	Password string  `gorm:"not null" json:"password"`
	Avatar   string  `json:"avatar"`
	Role     string  `gorm:"default:'user';check:role in ('user','admin')" json:"role"`
	Status   string  `gorm:"default:'active';check:role in ('active','disabled')" json:"status"`
	Orders   []Order `json:"orders"`
}
