package net

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type LoginResData struct {
	Token string `json:"token"`
}

type CategoryInfoReqData struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ProductCount int    `json:"productCount"`
}

func NewRes(code int, data any, message string) Response {
	return Response{Code: code, Data: data, Message: message}
}
