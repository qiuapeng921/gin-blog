package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(c *gin.Context) {
	conn, err := opOpen(c.Writer, c.Request)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("用户下线", err.Error())
			onClone()
			break
		}
		log.Println(string(msg))
		if string(msg) == "PING" {
			err = conn.WriteMessage(msgType, []byte("DONG"))
			if err != nil {
				break
			}
		} else {
			// todo：业务操作
			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				break
			}
		}
	}
}

func opOpen(response http.ResponseWriter, request *http.Request) (conn *websocket.Conn, err error) {
	conn, err = wsUpGrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println("websocket upgrade err:", err.Error())
		return
	}
	_ = conn.WriteMessage(websocket.TextMessage, []byte("welcome"))
	go sendTime(conn)
	return
}

func onClone() {

}

func sendTime(conn *websocket.Conn) {
	for {
		time.Sleep(1 * time.Second)
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		_ = conn.WriteMessage(websocket.TextMessage, []byte(nowTime))
	}
}
