package common

type Response[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

type PageData[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

type PageResponse[T any] struct {
	Response[T]
	Data PageData[T] `json:"data"`
}
