package controller

import (
	"gin-blog/app/consts"
	"gin-blog/app/models/articles"
	"gin-blog/app/models/categorys"
	"gin-blog/app/request"
	"gin-blog/app/service"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Articles(c *gin.Context) {
	var indexRequest request.IndexRequest
	if err := c.ShouldBindQuery(&indexRequest); err != nil {
		response.Context(c).Error(consts.ERROR)
		return
	}
	var article service.ArticleService
	result := article.GetList(indexRequest)
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

func Markdown(c *gin.Context) {
	var category []categorys.Entity
	grom.GetOrm().Find(&category)
	response.Context(c).View("markdown", gin.H{
		"category": category,
	})
}

func SaveArticle(c *gin.Context) {
	var form request.ArticleRequest
	if err := c.ShouldBind(&form); err != nil {
		response.Context(c).Error(consts.ERROR)
		return
	}
	var articleModel articles.Entity
	articleModel.Title = form.Title
	articleModel.Content = form.Content
	articleModel.CategoryId = form.Category
	articleModel.CreateTime = time.Now().Unix()
	articleModel.UpdateTime = time.Now().Unix()
	grom.GetOrm().Save(&articleModel)
	c.Redirect(http.StatusMovedPermanently, "/")
}
