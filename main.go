package main

import (
	"fmt"
	"gin-blog/app/crontab"
	"gin-blog/helpers/logging"
	"gin-blog/helpers/pool/gredis"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/system"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func init() {
	system.SetUp()
	logging.Setup()
	grom.SetUp()
	gredis.SetupRedis()
	crontab.InitCronTab()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func main() {
	gin.SetMode(os.Getenv("APP_ENV"))

	engine := gin.New()
	// 加载模板
	engine.LoadHTMLGlob("templates/*")

	// 设置路由
	routers.SetupRouter(engine)
	endPoint := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Println("|-----------------------------------|")
	log.Println("|             gin-blog              |")
	log.Println("|-----------------------------------|")
	log.Println("|  Go Http Server Start Successful  |")
	log.Println("|    Port" + endPoint + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	log.Println("|-----------------------------------|")

	_ = server.ListenAndServe()
}
