package controller

import (
	"gin-blog/app/models/admins"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
	"log"
)

func Index(c *gin.Context) {
	result := admins.GetOne(1)
	log.Println(result)
	response.Context(c).View("index", gin.H{
		"name": "GinFrame",
	})
}
