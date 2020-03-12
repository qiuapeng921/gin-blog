package system

import (
	"gin-blog/app/crontab"
	"gin-blog/helpers/logging"
	"gin-blog/helpers/pool/gredis"
	"gin-blog/helpers/pool/grom"
	"github.com/joho/godotenv"
	"log"
)

func SetUp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	logging.Setup()
	grom.SetUp()
	gredis.SetupRedis()
	crontab.InitCronTab()
}
