package controller

import (
	"go-web-demo/model"
	"go-web-demo/services"
	"go-web-demo/utils"

	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	var article model.Article
	c.Bind(&article)
	if articleId, err := services.AddArticle(&article); err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatResponse(0, articleId, ""))
	}
}

func GetArticleList(c *gin.Context) {
	var query model.ArticleQuery
	c.Bind(&query)
	articles, count, err := services.GetArticleList(&query)
	if err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatPageResponse(articles, count))
	}
}
