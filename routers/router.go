package routers

import (
	"gin-blog/app/controller"
	"gin-blog/app/middleware"
	"gin-blog/app/socket"
	"gin-blog/helpers/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.Use(middleware.JWT())

	router.GET("/", controller.Index)
	router.GET("/ws", socket.WsHandler)

	router.NoRoute(func(context *gin.Context) {
		response.Context(context).View("index", gin.H{
			"name": "GinFrame",
		})
	})
}
