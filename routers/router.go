package routers

import (
	"gin-blog/app/http/controller/api"
	"gin-blog/app/http/middleware/apiMiddleware"
	"gin-blog/app/socket"
	_ "gin-blog/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(apiMiddleware.JWT())

	router.Static("/static", "./resources")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	SetupRouterApi(router)

	router.GET("/", api.Index)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ws", socket.WsHandler)
}
