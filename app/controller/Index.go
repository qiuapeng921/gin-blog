package controller

import (
	"gin-blog/app/consts"
	"gin-blog/app/request"
	"gin-blog/app/service"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var indexRequest request.IndexRequest
	if err := c.ShouldBind(&indexRequest); err != nil {
		response.Context(c).Error(consts.ERROR)
	}
	var article service.ArticleService
	result := article.GetList(indexRequest)
	response.Context(c).View("index", gin.H{
		"title":   "首页",
		"article": result,
	})
}
