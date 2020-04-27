package elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"os"
)

var	conn *elastic.Client

func SetupElastic() {
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTIC_HOST")))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = client.Ping(os.Getenv("ELASTIC_HOST")).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	conn = client
	fmt.Println("elastic连接成功")
}

func GetConn() *elastic.Client {
	return conn
}
