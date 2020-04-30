package system

import (
	"gin-blog/app/crontab"
	"gin-blog/helpers/logging"
	"gin-blog/helpers/pool/gredis"
	"gin-blog/helpers/pool/grom"
	"gin-blog/helpers/pool/mongo"
	"gin-blog/helpers/pool/rabbit"
	"github.com/joho/godotenv"
	"log"
)

func SetUp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	logging.Setup()
	crontab.InitCronTab()
	grom.SetUpOrm()
	gredis.SetupRedis()
	mongo.SetupMongo()
	//elastic.SetupElastic()
	rabbit.SetupRabbitMq()
}