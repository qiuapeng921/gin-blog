package routers

import (
	"gin-blog/app/controller/api"
	"github.com/gin-gonic/gin"
)

func SetupRouterApi(router *gin.Engine)  {

	user := router.Group("/user")
	{
		user.GET("/list", api.UserList)
		user.GET("/info", api.UserInfo)
	}
}
