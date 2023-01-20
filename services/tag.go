package services

import (
	"fmt"
	"go-web-demo/model"
	"go-web-demo/utils"
)

func init() {
	fmt.Println("init tag service")
	utils.Db.AutoMigrate(&model.Tag{})
}

func GetTagList(query *model.TagQuery) ([]model.Tag, int64, error) {
	fmt.Println(query)
	var sql = utils.Db.Table("tags")
	var tags []model.Tag
	if query.PageNum == 0 {
		query.PageNum = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 15
	}
	if query.Name != "" {
		sql = sql.Where("name like ?", "%"+query.Name+"%")
	}
	var count int64
	result := sql.Count(&count)
	if result.Error != nil {
		return nil, count, result.Error
	}
	result = sql.Limit(query.PageSize).Offset((query.PageNum - 1) * query.PageSize).Find(&tags)
	return tags, count, result.Error
}

func AddTag(tag *model.Tag) (int64, error) {
	result := utils.Db.Create(tag)
	return tag.ID, result.Error
}
