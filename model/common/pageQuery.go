package common

type PageQuery struct {
	PageNum  int `json:"pageNum" default:"1"`
	PageSize int `json:"pageSize" default:"15"`
}
