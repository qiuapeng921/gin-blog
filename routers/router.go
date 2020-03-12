package routers

import (
	"gin-blog/app/controller/admin"
	"gin-blog/app/middleware"
	"gin-blog/app/socket"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.JWT())

	router.GET("/", admin.Index)
	router.GET("/ws", socket.WsHandler)
}
