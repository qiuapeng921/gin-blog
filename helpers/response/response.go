package response

import (
	"gin-blog/app/consts"
	"github.com/gin-gonic/gin"
)

// Wrapper include context
type Wrapper struct {
	*gin.Context
}

// WrapContext
func WrapContext(c *gin.Context) *Wrapper {
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
