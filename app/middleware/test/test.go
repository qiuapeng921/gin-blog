package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("请求地址为：", c.Request.URL)
	}
}
