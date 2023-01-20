package services

import (
	"fmt"
	"go-web-demo/model"
	"go-web-demo/utils"
)

func init() {
	fmt.Println("init article service")
	utils.Db.AutoMigrate(&model.Article{})
}

func AddArticle(article *model.Article) (int64, error) {
	for _, value := range article.TagIds {
		tag := model.Tag{}
		tag.ID = value
		article.Tags = append(article.Tags, tag)
	}
	result := utils.Db.Create(article)
	return article.ID, result.Error
}

func GetArticleList(params *model.ArticleQuery) (articles []model.Article, count int64, err error) {
	var sql = utils.Db.Model(&model.Article{})
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 15
	}

	if params.Title != "" {
		sql = sql.Where("title like ?", "%"+params.Title+"%")
	}

	if len(params.Tags) > 0 {
		sql = sql.Where("id in (?)", utils.Db.Table("article_tag").Select("article_id").Where("tag_id in ?", params.Tags))
	}

	sql.Preload("Tags")
	result := sql.Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	result = sql.Find(&articles)

	return articles, count, result.Error
}
