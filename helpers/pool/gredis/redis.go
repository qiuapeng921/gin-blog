package gredis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"strconv"
)

var conn *redis.Client

func SetupRedis() {
	Address := os.Getenv("REDIS_HOST")
	PassWord := os.Getenv("REDIS_PASSWORD")
	Idle, _ := strconv.Atoi(os.Getenv("MAX_IDLE"))
	Active, _ := strconv.Atoi(os.Getenv("MAX_ACTIVE"))
	conn = redis.NewClient(&redis.Options{
		Addr:         Address,
		Password:     PassWord,
		DB:           0,
		PoolSize:     Idle,
		MinIdleConns: Active,
	})
	pong, err := conn.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功", pong)
}

func GetConn() *redis.Client {
	return conn
}
