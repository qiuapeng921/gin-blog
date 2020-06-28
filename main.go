package main

import (
	"fmt"
	"gin-blog/app/crontab"
	"gin-blog/helpers/logging"
	"gin-blog/helpers/pool/gredis"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/pool/rabbit"
	"gin-blog/helpers/templates"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	logging.Setup()
	crontab.InitCronTab()
	grom.SetUpOrm()
	gredis.SetupRedis()
	//mongo.SetupMongo()
	//elastic.SetupElastic()
	rabbit.SetupRabbitMq()
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
	log.Printf("weclome use gin,地址:http://127.0.0.1:%s \n", os.Getenv("HTTP_PORT"))
	_ = server.ListenAndServe()
}
