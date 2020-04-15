package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Println("requestUrl：", c.Request.URL)
	}
}
