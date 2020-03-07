package api

import (
	"gin-blog/app/consts"
	"gin-blog/app/service/user_service"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description 获取所有用户
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /user/list [get]
func UserList(context *gin.Context) {
	app := response.Gin{Context: context}
	result := user_service.GetUserAll()
	app.Response(http.StatusOK, consts.SUCCESS, result)
}

// @Description 获取指定ID记录
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /user/info [get]
func UserInfo(context *gin.Context) {
	app := response.Gin{Context: context}
	id := context.GetInt("id")
	result := user_service.GetOne(id)
	app.Response(http.StatusOK, consts.SUCCESS, result)
}
