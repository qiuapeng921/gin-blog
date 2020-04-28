package routers

import (
	"gin-blog/app/controller"
	"gin-blog/app/middleware"
	"gin-blog/app/socket"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine) {
	router.Use(middleware.JWT())

	router.GET("/", controller.Index)
	router.GET("/ws", socket.WsHandler)

	router.GET("/markdown", controller.Markdown)
	router.GET("/articles", controller.Articles)
	router.GET("/articleInfo", controller.ArticleInfo)
	router.POST("/saveArticle", controller.SaveArticle)

	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/")
	})
}
