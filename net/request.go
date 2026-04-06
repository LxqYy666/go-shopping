package net

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddCategoryReq struct {
	Name string `json:"name"`
}

type AddProductReq struct {
	CategoryID uint    `json:"category_id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Stock      int     `json:"stock"`
	ImageUrl   string  `json:"image_url"`
}

type UpdateProductReq struct {
	CategoryID *uint    `json:"category_id,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Desc       *string  `json:"desc,omitempty"`
	Price      *float32 `json:"price,omitempty"`
	Stock      *int     `json:"stock,omitempty"`
	ImageUrl   *string  `json:"image_url,omitempty"`
	Status     *string  `json:"status,omitempty"`
}

type AddToCartReq struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type UpdateCartReq struct {
	Quantity int `json:"quantity"`
}

type CreateOrderReq struct {
	ReceiverAddr  string `json:"receiver_addr"`
	ReceiverName  string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	Remark        string `json:"remark"`
}

type UpdateOrderReq struct {
	Status *string `json:"status,omitempty"`
}
