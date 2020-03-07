package api

import (
	"github.com/gin-gonic/gin"
)

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /user/list [get]
func UserList(c *gin.Context) {
	c.String(200, "UserList")
}
