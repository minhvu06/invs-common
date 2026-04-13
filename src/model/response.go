package model

type BaseResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func Success[T any](data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		Code:    0,
		Message: "Thành công",
		Data:    data,
	}
}

func Error[T any](code int, msg string) *BaseResponse[T] {
	return &BaseResponse[T]{
		Code:    code,
		Message: msg,
		Data:    *new(T),
	}
}
