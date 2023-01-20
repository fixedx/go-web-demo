package router

import (
	"go-web-demo/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	r.GET("/users", controller.GetUserList)
	r.GET("/user/:id", controller.GetUserById)
}
