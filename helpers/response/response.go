package response

import (
	"fmt"
	"gin-blog/app/consts"
	"github.com/gin-gonic/gin"
)

type Wrapper struct {
	*gin.Context
}

func Context(c *gin.Context) *Wrapper {
	return &Wrapper{c}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response setting gin.JSON
func (wrapper *Wrapper) Response(httpCode, errCode int, data interface{}) {
	wrapper.JSON(httpCode, Response{
		Code:    errCode,
		Message: consts.GetMsg(errCode),
		Data:    data,
	})
	return
}

func (wrapper *Wrapper) View(name string, data interface{}) {
	wrapper.HTML(200, fmt.Sprintf("%s.html", name), data)
	return
}
