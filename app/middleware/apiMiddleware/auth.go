package apiMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("requestUrl：", c.Request.URL)
	}
}
