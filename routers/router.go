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

	router.Static("/static", "./public")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.GET("/", admin.Index)
	router.GET("/ws", socket.WsHandler)
}
