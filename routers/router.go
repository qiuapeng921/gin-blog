package routers

import (
	"gin-blog/app/controller/api"
	"gin-blog/app/middleware/apiMiddleware"
	_ "gin-blog/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(apiMiddleware.JWT())

	router.GET("/", api.Index)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/user")
	{
		user.GET("/list", api.UserList)
		user.GET("/info", api.UserInfo)
	}

	return router
}
