package common

import (
	"GMS/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	StatusSuccess           = 0   // 成功
	StatusVersionLow        = 100 // 版本太旧，请升级
	StatusUserNotExist      = 101 // 用户不存在
	StatusInvalidPassword   = 102 // 用户密码不匹配
	StatusUserHasRegistered = 103 // 用户已注册
	StatusInvalidToken      = 122 // 无效TOKEN
	StatusJsonError         = 141 // body中的json格式错误
	StatusInvalidParam      = 142 // 部分参数为空
	StatusServerError       = 400 // 服务器内部错误
	StatusMethodNotAllowed  = 405 // 无效method
)

const (
	MsgSuccess           = "success"
	MsgVersionLow        = "invalid version"
	MsgUserNotExist      = "user not exist"
	MsgInvalidPassword   = "invalid password"
	UserHasRegisteredMsg = "user hs registered"
	MsgServerError       = "server error"
	MsgInvalidToken      = "invalid token"
	MsgMethodNotAllowed  = "method not allowed"
)

//ResponseModel 响应数据,带参
type ResponseModel struct {
	Result    bool        `json:"result"`
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

// 响应成功(带返回值)
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Result: true, ErrorCode: StatusSuccess, Message: MsgSuccess, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功(不带返回值)
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModel{Result: true, ErrorCode: StatusSuccess, Message: MsgSuccess}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFailCode(c *gin.Context, code int, msg string) {
	ret := ResponseModel{Result: false, ErrorCode: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

//响应失败，带返回值
func ResFailData(c *gin.Context, code int, msg string, v interface{}) {
	ret := ResponseModel{Result: false, ErrorCode: code, Message: msg, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 传参异常
func ResParamFail(c *gin.Context) {
	ret := ResponseModel{Result: false, ErrorCode: StatusInvalidParam, Message: MsgSuccess}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	logger.InfoKV("ResJSON", logger.KV{"status": status, "val": v})
	c.JSON(status, v)
	c.Abort()
}
