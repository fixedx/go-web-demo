package router

import (
	"go-web-demo/controller"

	"github.com/gin-gonic/gin"
)

func TagRouter(r *gin.Engine) {
	r.POST("/tag", controller.AddTag)
	r.GET("/tags", controller.GetTagList)
}
