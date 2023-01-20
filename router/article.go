package router

import (
	"go-web-demo/controller"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.Engine) {
	r.POST("/article", controller.AddArticle)
	r.GET("/articles", controller.GetArticleList)
}
