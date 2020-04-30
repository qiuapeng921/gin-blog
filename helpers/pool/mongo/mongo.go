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

var collection *mongo.Database

//初始化
func SetupMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	maxPool, _ := strconv.Atoi(os.Getenv("MONGO_MAX_POOL"))
	//[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	//uri := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO_HOST"))
	//uri := fmt.Sprintf("mongodb://root:123456@192.168.0.163/admin")
	uri := fmt.Sprintf("mongodb://%s:%s@%s/admin",  os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_HOST"))
	var err error
	// 连接数据库
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(uint64(maxPool)))
	if err != nil {
		panic(err.Error())
	}
	// 判断服务是不是可用
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}
	database := os.Getenv("MONGO_DATABASE")
	collection = client.Database(database)
	fmt.Println("mongodb连接成功")
}

func GetMongoDb() *mongo.Client {
	return client
}

func GetMongoCollection(table string) *mongo.Collection {
	return collection.Collection(table)
}

func Create(table string, data interface{}) (*mongo.InsertOneResult, error) {
	fmt.Println(data)
	return GetMongoCollection(table).InsertOne(context.TODO(), data)
}

func CreateMany(table string, data []interface{}) (*mongo.InsertManyResult, error) {
	return GetMongoCollection(table).InsertMany(context.TODO(), data)
}

func Update(table string, filter, update interface{}) (*mongo.UpdateResult, error) {
	return GetMongoCollection(table).UpdateOne(context.TODO(), filter, update)
}

func FindOne(table string, filter interface{}) *mongo.SingleResult {
	return GetMongoCollection(table).FindOne(context.TODO(), filter)
}

func Delete(table string, filter interface{}) (*mongo.DeleteResult, error) {
	return GetMongoCollection(table).DeleteOne(context.TODO(), filter)
}