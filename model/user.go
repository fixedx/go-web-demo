package model

type User struct {
	Id      int64 `gorm:"primaryKey"`
	Name    string
	Address string
	Age     int
	Email   string
}

type UserQuery struct {
	Name   string `form:"name"`
	MaxAge int    `form:"maxAge"`
	MinAge int    `form:"minAge"`
}
