package model

import "go-web-demo/model/common"

type Tag struct {
	common.CommonDbModel
	Name string `gorm:"type:varchar(50);not null" json:"name"`
}

type TagQuery struct {
	common.PageQuery
	Name string `json:"name" form:"name"`
}
