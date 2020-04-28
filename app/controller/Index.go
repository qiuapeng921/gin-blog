package controller

import (
	"fmt"
	"gin-blog/app/consts"
	"gin-blog/app/request"
	"gin-blog/app/service"
	"gin-blog/helpers/pool/elastic"
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
	for _, v := range result {
		createResult, err := elastic.Create("article", "text", v)
		if err != nil {
			fmt.Println("es 数据添加失败", err.Error())
		}
		fmt.Println("es create successful", createResult.Id)
	}
	response.Context(c).View("index", gin.H{
		"title":   "首页",
		"article": result,
	})
}
