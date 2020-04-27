package controller

import (
	"gin-blog/app/service"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var article service.ArticleService
	result := article.GetList()
	response.Context(c).View("index", gin.H{
		"title":   "首页",
		"article": result,
	})
}
