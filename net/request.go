package net

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func NewRes(code int, data any, message string) Response {
	return Response{Code: code, Data: data, Message: message}
}
