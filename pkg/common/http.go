package common

import (
	"GMS/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	StatusSuccess           = 0   // 成功
	//100 - 200 sys
	StatusVersionLow        = 100 // 版本太旧，请升级
	StatusMethodNotAllowed  = 101 // 无效method
	StatusServerError       = 102 // 服务器内部错误
	StatusJsonError         = 103 // body中的json格式错误
	StatusInvalidParam      = 104 // 部分参数为空

	//200 - 300 user
	StatusUserNotExist      = 201 // 用户不存在
	StatusInvalidPassword   = 202 // 用户密码不匹配
	StatusUserHasRegistered = 203 // 用户已注册
	StatusInvalidToken      = 204 // 无效TOKEN

	//300 - 400 role
	StatusRoleNotExist      = 301 // 角色不存在
)

const (
	MsgSuccess           = "success"

	//sys
	MsgVersionLow        = "invalid version"
	MsgMethodNotAllowed  = "method not allowed"
	MsgServerError       = "server error"
	MsgInvalidParam      = "invalid param"

	//user
	MsgUserNotExist      = "user not exist"
	MsgInvalidPassword   = "invalid password"
	MsgUserHasRegistered = "user hs registered"
	MsgInvalidToken      = "invalid token"

	//role
	MsgRoleNotExist      = "role not exist"
)

//ResponseModel 响应数据,带参
type ResponseModel struct {
	ErrorCode int         `json:"code"`
	Message   string      `json:"detail"`
	Data      interface{} `json:"data,omitempty"`
}

// 响应成功(带返回值)
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{ErrorCode: StatusSuccess, Message: MsgSuccess, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功(不带返回值)
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModel{ErrorCode: StatusSuccess, Message: MsgSuccess}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFailCode(c *gin.Context, code int, msg string) {
	ret := ResponseModel{ErrorCode: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

//响应失败，带返回值
func ResFailData(c *gin.Context, code int, msg string, v interface{}) {
	ret := ResponseModel{ErrorCode: code, Message: msg, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 传参异常
func ResParamFail(c *gin.Context) {
	ret := ResponseModel{ErrorCode: StatusInvalidParam, Message: MsgSuccess}
	ResJSON(c, http.StatusOK, &ret)
}

func ResFailError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, err)
	c.Abort()
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	logger.InfoKV("ResJSON", logger.KV{"status": status, "val": v})
	c.JSON(status, v)
	c.Abort()
}
