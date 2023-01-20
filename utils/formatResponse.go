package utils

import (
	"go-web-demo/model/common"
)

func FormatResponse[T any](code int, data T, msg string) common.Response[T] {
	return common.Response[T]{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func FormatSuccessResponse[T any](data T) common.Response[T] {
	return FormatResponse(0, data, "")
}

func FormatPageResponse[T any](data []T, total int64) common.Response[common.PageData[T]] {
	var pageData = common.PageData[T]{
		List:  data,
		Total: total,
	}
	return FormatResponse(0, pageData, "")
}
