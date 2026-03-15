package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        uint    `json:"user_id"`
	TotalAmount   float32 `gorm:"not null" json:"total_amount"`
	Status        string  `gorm:"default:'pending';check:status in ('pending','paid','shipped','completed')"`
	ReceiverAddr  string  `gorm:"not null" json:"receiver_addr"`
	ReceiverName  string  `gorm:"not null" json:"receiver_name"`
	ReceiverPhone string  `gorm:"not null" json:"receiver_phone"`
	Remark        string  `json:"remark"`

	User  *User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Items []OrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `gorm:"not null" json:"quantity"`

	// 关联
	Order   *Order   `json:"-" gorm:"foreignKey:OrderID"`
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
