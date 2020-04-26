package elastic

import (
	"context"
	"github.com/olivere/elastic"
	"os"
)

type Elastic struct {
	conn *elastic.Client
}

func SetupElastic() {
	conn, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTIC_HOST")))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = conn.Ping(os.Getenv("ELASTIC_HOST")).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	var elastic Elastic
	elastic.conn = conn
}

func (e *Elastic) GetConn() *elastic.Client {
	return e.conn
}
