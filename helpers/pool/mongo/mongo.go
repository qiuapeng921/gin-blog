package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var db *Database

//初始化
func Setup() {
	db = &Database{
		Mongo: setConnect(),
	}
}

// 连接设置
func setConnect() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO_HOST"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	maxPool, _ := strconv.Atoi(os.Getenv("MONGO_MAX_POOL"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(uint64(maxPool)))
	if err != nil {
		panic(err.Error())
	} // 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func GetMongoDb() *Database {
	return db
}
