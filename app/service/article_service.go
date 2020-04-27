package service

import (
	"gin-blog/app/models/articles"
	"gin-blog/helpers/pool/grom"
)

type ArticleService struct {

}

func (a *ArticleService) GetList() (article []articles.Entity) {
	grom.GetConn().Find(&article)
	return article
}

func (a *ArticleService) GetOne(id int) (article articles.Entity) {
	grom.GetConn().Where("id = ?", id).Find(&article)
	return article
}