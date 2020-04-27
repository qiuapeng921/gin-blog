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

var client *mongo.Client

//初始化
func SetupMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	maxPool, _ := strconv.Atoi(os.Getenv("MONGO_MAX_POOL"))
	uri := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO_HOST"))
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(uint64(maxPool)))

	if err != nil {
		panic(err.Error())
	} // 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("mongodb连接成功")
}

func GetMongoDb() *mongo.Client {
	return client
}
