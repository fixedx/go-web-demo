package services

import (
	"fmt"
	"go-web-demo/model"
	"go-web-demo/utils"
)

func init() {
	fmt.Println("init user services")
	utils.Db.AutoMigrate(&model.User{})
}

func GetUserList(query *model.UserQuery) []model.User {
	var users = []model.User{}
	var sql = utils.Db.Table("user")
	if query.Name != "" {
		sql = sql.Where("name like ? ", query.Name+"%")
	}
	if query.MaxAge != 0 && query.MinAge != 0 {
		sql = sql.Where("age between ? and ?", query.MinAge, query.MaxAge)
	} else if query.MinAge != 0 {
		sql = sql.Where("age >= ?", query.MinAge)
	} else if query.MaxAge != 0 {
		sql = sql.Where("age <= ?", query.MaxAge)
	}
	sql.Find(&users)
	return users
}

func GetUser(user *model.User) model.User {
	var res model.User
	utils.Db.Table("user").Where(&user).Find(&res)
	return res
}
