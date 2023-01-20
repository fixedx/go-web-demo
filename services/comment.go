package services

import (
	"fmt"
	"go-web-demo/model"
	"go-web-demo/utils"
)

func init() {
	fmt.Println("init comment services")
	utils.Db.AutoMigrate(&model.Comment{})
}

func GetCommentList(params *model.CommentQuery) (comments []model.Comment, count int64, err error) {
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 15
	}
	query := utils.Db.Model(&model.Comment{}).Where("article_id = ? and parent_comment_id = 0", params.ArticleId).Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize)
	result := query.Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	result = query.Find(&comments)
	return comments, count, result.Error
}

func AddComment(comment *model.Comment) (int64, error) {
	result := utils.Db.Create(comment)
	return comment.ID, result.Error
}
