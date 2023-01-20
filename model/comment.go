package model

import "go-web-demo/model/common"

type Comment struct {
	common.CommonDbModel
	Content         string `gorm:"type:varchar(255);not null" json:"content"`
	ParentCommentId int64  `json:"parentCommentId"`
	ArticleId       int64  `json:"articleId"`
}

type CommentQuery struct {
	common.PageQuery
	ArticleId int64 `json:"articleId" form:"articleId"`
}
