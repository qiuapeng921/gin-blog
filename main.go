package main

import (
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
	routersHandler := routers.InitRouter()
	routersHandler.LoadHTMLGlob("templates/*")

	endPoint := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersHandler,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start response server listening %s", endPoint)
	_ = server.ListenAndServe()
}
