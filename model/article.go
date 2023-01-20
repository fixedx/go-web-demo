package model

import "go-web-demo/model/common"

type Article struct {
	common.CommonDbModel
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Content  string `gorm:"type:text;not null" json:"content"`
	Author   string `gorm:"default:fixedx" json:"author"`
	Stars    int64  `json:"stars"`
	Comments []Comment
	Tags     []Tag   `gorm:"many2many:article_tag" json:"tags"`
	TagIds   []int64 `gorm:"-"`
}

type ArticleQuery struct {
	common.PageQuery
	Title string  `json:"title" form:"title"`
	Tags  []int64 `json:"tags" form:"tags"`
}
