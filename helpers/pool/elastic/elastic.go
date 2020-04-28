package elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"os"
)

var client *elastic.Client

func SetupElastic() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(os.Getenv("ELASTIC_HOST")))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = client.Ping(os.Getenv("ELASTIC_HOST")).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("elastic连接成功")
}

func GetConn() *elastic.Client {
	return client
}

//创建
func Create(Database, Type string, Data interface{}) (*elastic.IndexResponse, error) {
	return client.Index().
		Index(Database).
		Type(Type).
		BodyJson(Data).
		Do(context.Background())
}

//删除
func Delete(Database, Type string) (*elastic.DeleteResponse, error) {
	return client.Delete().Index(Database).
		Type(Type).
		Id("1").
		Do(context.Background())
}

//修改
func Update(Database, Type string, Data interface{}) (*elastic.UpdateResponse, error) {
	return client.Update().
		Index(Database).
		Type(Type).
		Id("2").
		Doc(Data).
		Do(context.Background())
}

//查找
func Query(Database, Type string) (*elastic.GetResult, error) {
	//通过id查找
	return client.Get().Index(Database).Type(Type).Id("2").Do(context.Background())
}