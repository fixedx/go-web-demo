package router

import (
	"go-web-demo/controller"

	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.Engine) {
	r.POST("/comment", controller.AddComment)
	r.GET("/comments", controller.GetCommentListByArticleId)
}
