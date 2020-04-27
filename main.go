package main

import (
	"fmt"
	"gin-blog/helpers/system"
	"gin-blog/helpers/templates"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	system.SetUp()
}

func main() {
	gin.SetMode(os.Getenv("APP_ENV"))

	engine := gin.Default()
	// 加载模板和资源文件
	templates.InitTemplate(engine)
	// 设置路由
	routers.SetupRouter(engine)

	endPoint := fmt.Sprintf("%s:%s", os.Getenv("HTTP_ADDRESS"), os.Getenv("HTTP_PORT"))

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Go Http Server Start Successful")
	log.Println("Port:" + endPoint + ",Pid:" + fmt.Sprintf("%d", os.Getpid()))

	_ = server.ListenAndServe()
}
