package controller

import (
	"go-web-demo/model"
	"go-web-demo/services"
	"go-web-demo/utils"

	"github.com/gin-gonic/gin"
)

func AddTag(c *gin.Context) {
	var tag model.Tag
	c.Bind(&tag)
	if tagId, err := services.AddTag(&tag); err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatSuccessResponse(tagId))
	}
}

func GetTagList(c *gin.Context) {
	var query model.TagQuery
	c.ShouldBindQuery(&query)
	if result, count, err := services.GetTagList(&query); err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatPageResponse(result, count))
	}

}
