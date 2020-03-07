package main

import (
	"context"
	"fmt"
	"gin-blog/helpers/logging"
	"gin-blog/helpers/pool/gredis"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/system"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	system.SetUp()
	grom.SetUp()
	_ = gredis.SetupRedis()
	logging.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func main() {
	gin.SetMode(os.Getenv("APP_ENV"))

	engine := gin.New()

	// 设置路由
	routers.SetupRouter(engine)

	engine.LoadHTMLGlob("templates/*")

	endPoint := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-gin-api             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port" + endPoint + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
