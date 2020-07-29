package cronfunc

import (
	"fmt"
	"gin-blog/helpers/pool/rabbit"
	"github.com/rs/xid"
)

func TestPushMessage() {
	id := xid.New().String()
	err := rabbit.Publish("test", "test", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("推送", id)
}
