package routers

import (
	"gin-blog/app/controller/api"
	"gin-blog/app/middleware/apiMiddleware"
	_ "gin-blog/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter(router *gin.Engine) {

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(apiMiddleware.JWT())

	SetupRouterApi(router)

	router.GET("/", api.Index)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
