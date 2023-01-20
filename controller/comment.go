package controller

import (
	"fmt"
	"go-web-demo/model"
	"go-web-demo/services"
	"go-web-demo/utils"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var comment model.Comment
	c.Bind(&comment)
	fmt.Println(comment)
	commentId, err := services.AddComment(&comment)
	if err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatResponse(0, commentId, ""))
	}
}

func GetCommentListByArticleId(c *gin.Context) {
	var params model.CommentQuery
	c.Bind(&params)
	comments, count, err := services.GetCommentList(&params)
	if err != nil {
		c.JSON(200, utils.FormatResponse[any](0, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatPageResponse(comments, count))
	}
}
