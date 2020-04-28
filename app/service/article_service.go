package service

import (
	"gin-blog/app/models/articles"
	"gin-blog/app/request"
	"gin-blog/helpers/pool/grom"
)

type ArticleService struct {
}

func (a *ArticleService) GetList(request request.IndexRequest) (article []articles.Entity) {
	orm := grom.GetConn()
	if request.Id > 0 {
		orm = orm.Where("category_id = ?", request.Id)
	}
	if request.Page > 0 {
		orm = orm.Offset(request.Page)
	}
	orm.Order("id desc").Find(&article)
	return article
}

func (a *ArticleService) GetOne(id int) (article articles.Entity) {
	grom.GetConn().Where("id = ?", id).Find(&article)
	return article
}
