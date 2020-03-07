package api

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	Example
	c.HTML(200, "index.tmpl", gin.H{
		"name": "GinFrame",
	})
}