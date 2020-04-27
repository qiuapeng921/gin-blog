package controller

import (
	"gin-blog/app/consts"
	"gin-blog/app/models/articles"
	"gin-blog/app/request"
	"gin-blog/app/service"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Articles(c *gin.Context) {
	var article service.ArticleService
	result := article.GetList()
	response.Context(c).View("list", gin.H{"article": result})
}

func ArticleInfo(c *gin.Context) {
	params := c.Query("id")
	id, _ := strconv.Atoi(params)
	var articleService service.ArticleService
	article := articleService.GetOne(id)
	response.Context(c).View("info", gin.H{
		"article": article,
	})
}

func SaveArticle(c *gin.Context) {
	var form request.ArticleRequest
	err := c.ShouldBind(&form)
	if err != nil {
		response.Context(c).Error(consts.ERROR)
		return
	}
	var articleModel articles.Entity
	articleModel.Title = form.Title
	articleModel.Content = form.Content
	grom.GetConn().Save(&articleModel)
	c.Redirect(http.StatusMovedPermanently, "/markdown")
}
