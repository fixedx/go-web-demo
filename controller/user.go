package controller

import (
	"fmt"
	"go-web-demo/model"
	"go-web-demo/services"
	"go-web-demo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	var query model.UserQuery
	c.Bind(&query)
	users := services.GetUserList(&query)
	c.JSON(200, utils.FormatSuccessResponse(users))
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
		user := model.User{Id: id}
		user = services.GetUser(&user)
		c.JSON(200, utils.FormatSuccessResponse(user))
	} else {
		fmt.Println(err)
		c.JSON(400, utils.FormatResponse[any](400, nil, "参数错误"))
	}
}
