package net

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type LoginResData struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type CategoryInfoReqData struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ProductCount int    `json:"productCount"`
}

type ProductInfoReqData struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	CategoryId uint    `json:"category_id"`
	Price      float32 `json:"price"`
	Stock      int     `json:"stock"`
	ImageUrl   string  `json:"image_url"`
	SoldCount  int     `json:"sold_count"`
	Status     string  `json:"status"`
}

type UserInfoReqData struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	OrdersCount int    `json:"orders_count"`
	CreatedAt   string `json:"created_at"`
}

type CartItemData struct {
	ID        uint            `json:"id"`
	ProductID uint            `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Product   ProductInfoReqData `json:"product"`
}

type OrderItemData struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float32 `json:"total_price"`
	Product    *ProductInfoReqData `json:"product,omitempty"`
}

type OrderData struct {
	ID            uint            `json:"id"`
	UserID        uint            `json:"user_id"`
	TotalAmount   float32         `json:"total_amount"`
	Status        string          `json:"status"`
	ReceiverAddr  string          `json:"receiver_addr"`
	ReceiverName  string          `json:"receiver_name"`
	ReceiverPhone string          `json:"receiver_phone"`
	Remark        string          `json:"remark"`
	CreatedAt     string          `json:"created_at"`
	User          *UserInfoReqData `json:"user,omitempty"`
	Items         []OrderItemData `json:"items,omitempty"`
}

func NewRes(code int, data any, message string) Response {
	return Response{Code: code, Data: data, Message: message}
}
