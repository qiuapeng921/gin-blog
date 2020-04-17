package controller

import (
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Context(c).View("markdown", gin.H{
		"name": "GinFrame",
	})
}
