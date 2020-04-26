package gredis

import (
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	conn *redis.Pool
}

func SetupRedis()  {
	Idle, _ := strconv.Atoi(os.Getenv("MAX_IDLE"))
	Active, _ := strconv.Atoi(os.Getenv("MAX_ACTIVE"))
	var RedisConn Redis
	RedisConn.conn = &redis.Pool{
		MaxIdle:     Idle,    //最大空闲数
		MaxActive:   Active,  //最大连接数据库连接数
		IdleTimeout: 200, // 超时
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			panic(err.Error())
		},
	}
}

func GetInstance() *Redis {
	return &Redis{}
}