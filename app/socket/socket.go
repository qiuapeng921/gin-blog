package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// Msg 消息结构体
type Msg struct {
	Type string      `json:"type" v:"type@required#消息类型不能为空"`
	Data interface{} `json:"data" v:""`
	From string      `json:"name" v:""`
}

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	var err error
	conn, err = wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade err:", err.Error())
		return
	}

	_ = conn.WriteMessage(websocket.TextMessage, []byte("welcome"))

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("用户下线", err.Error())
			break
		}
		log.Println(string(msg))
		if string(msg) == "PING" {
			err = conn.WriteMessage(msgType, []byte("DONG"))
			if err != nil {
				break
			}
			return
		}
		// todo：业务操作
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
	}
}
