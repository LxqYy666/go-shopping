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

func NewRes(code int, data any, message string) Response {
	return Response{Code: code, Data: data, Message: message}
}
