package controller

import (
	"fmt"
	"gin-blog/app/models/articles"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Context(c).View("index", gin.H{
		"name": "GinFrame",
	})
}

func Articles(c *gin.Context) {
	response.Context(c).View("list")
}

func ArticleInfo(c *gin.Context) {
	response.Context(c).View("info")
}

func SaveArticle(c *gin.Context) {
	params := c.PostForm("editor-markdown-doc")
	fmt.Println(params)
	var entity articles.Entity
	entity.Content = params
	db := grom.GetConn().Save(&entity)
	fmt.Println(db.RowsAffected)
}
