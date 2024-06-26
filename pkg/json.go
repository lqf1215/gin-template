package pkg

import (
	"gin-template/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	CodeOk       = 0  // 成功
	CodeErr      = -1 // 失败
	CodeErrToken = -2 // token相关的异常
	//CodeReject   = "2" // 拒绝
	//CodeTimeout  = "3" // 超时
)

// JSONResponse represents an HTTP response which contains a JSON body.
type JSONResponse struct {
	// HTTP status code.
	Code int `json:"code"`
	// JSON represents the JSON that should be serialized and sent to the client
	Data interface{} `json:"data"`
}

func OkResponse(c *gin.Context, data interface{}) {
	c.JSON(200, JSONResponse{
		Code: 0,
		Data: data,
	})
}

// MessageResponse returns a JSONResponse with a 'message' key containing the given text.
func MessageResponse(c *gin.Context, code int, msg, msgZh string) {
	global.Log.Warn(msgZh, zap.Int("code", code),
		zap.String("msg_zh", msgZh))
	//Log.Warnf("12312")
	c.JSON(code, JSONResponse{
		Code: code,
		Data: struct {
			Message   string `json:"message"`
			MessageZh string `json:"message_zh"`
		}{msg, msgZh},
	})
}
