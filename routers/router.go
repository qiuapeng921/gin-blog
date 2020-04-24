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

	router.GET("/articles", controller.Articles)
	router.GET("/articleInfo", controller.ArticleInfo)
	router.POST("/saveArticle", controller.SaveArticle)

	router.GET("markdown", func(context *gin.Context) {
		response.Context(context).View("markdown")
	})

	router.NoRoute(func(context *gin.Context) {
		context.Redirect(301, "/")
	})
}
