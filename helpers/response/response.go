package response

import (
	"gin-blog/app/consts"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Context *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.Context.JSON(httpCode, Response{
		Code: errCode,
		Msg:  consts.GetMsg(errCode),
		Data: data,
	})
	return
}
