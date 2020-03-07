package api

import (
	"fmt"
	"gin-blog/app/consts"
	"gin-blog/app/models"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Description 获取所有用户
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /user/list [get]
func UserList(c *gin.Context) {
	result, _ := models.GetUserALl()
	fmt.Println(result)
	c.JSON(200, gin.H{
		"data": result,
	})
}

// @Description 获取指定ID记录
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	userId, _ := strconv.Atoi(id)
	result, _ := models.GetOneById(userId)
	fmt.Println(result)
	response.WrapContext(c).Response(http.StatusOK, consts.SUCCESS, result)
}
