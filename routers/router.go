package routers

import (
	"errors"
	"gin-blog/app/controller"
	"gin-blog/app/middleware"
	"gin-blog/app/socket"
	"gin-blog/helpers/pool/gredis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/rs/xid"
	"math/rand"
	"net/http"
)

func SetupRouter(router *gin.Engine) {
	router.Use(middleware.RequestLog())

	router.GET("/", controller.Index)
	router.GET("/ws", socket.WsHandler)

	router.GET("/markdown", controller.Markdown)
	router.GET("/articles", controller.Articles)
	router.GET("/articleInfo", controller.ArticleInfo)
	router.POST("/saveArticle", controller.SaveArticle)

	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/")
	})

	router.GET("/order", func(context *gin.Context) {
		redisClient := gredis.GetRedis()
		err := redisClient.Watch(func(tx *redis.Tx) error {
			num := rand.Intn(5) + 1
			orderNumber, _ := tx.Get("order_number").Int()
			if num > orderNumber {
				return errors.New("已售空")
			}
			_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
				// 在这里给key设置最新值
				decr := pipe.DecrBy("order_number", int64(num))
				_, err := pipe.Exec()
				if err != nil {
					return err
				}
				if decr.Val() <= 0 {
					return errors.New("异常")
				}
				tx.HSetNX("userList", xid.New().String(), num)
				tx.IncrBy("buy_number", int64(num))
				return nil
			})
			if err != nil {
				return err
			}
			return nil
		}, "order_number")
		if err != nil {
			context.JSON(http.StatusOK, err.Error())
			return
		}
		user, _ := redisClient.Get("userList").Result()
		response := make(map[string]interface{}, 2)
		response["message"] = "下单成功"
		response["data"] = user
		context.JSON(http.StatusOK, response)
		return
	})
}
