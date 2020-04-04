package routers

import (
	"gin-blog/app/controller"
	"gin-blog/app/middleware"
	"gin-blog/app/socket"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.Use(middleware.JWT())

	router.GET("/", controller.Index)
	router.GET("/ws", socket.WsHandler)
}
