package middleware

import (
	"GMS/pkg/common"
	"GMS/pkg/jwt"
	"github.com/gin-gonic/gin"
)

const (
	TokenKey = "Access-Token"
	UserKey = "User-Info"
	ExpTimeout = 24*60*60 // token过期时间 second
)

func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		//获取token
		var token string
		if token = c.GetHeader(TokenKey); token == "" {
			common.ResFailCode(c, common.StatusInvalidToken, common.MsgInvalidToken)
			return
		}

		//获取userInfo
		var userInfo map[string]string
		var ok bool
		if userInfo, ok = jwt.ParseToken(token); !ok {
			common.ResFailCode(c, common.StatusInvalidToken, common.MsgInvalidToken)
			return
		}

		//数据是否为空
		if userInfo["uuid"] == "" || userInfo["userId"] == "" {
			common.ResFailCode(c, common.StatusInvalidToken, common.MsgInvalidToken)
			return
		}

		//查询redis key是否存在
		//if exist, err := cache.Exist(userInfo["userId"]); err != nil {
		//	common.ResFailCode(c, common.StatusServerError, common.MsgInvalidToken)
		//	return
		//} else if !exist {
		//	//数据不存在,token过期
		//	common.ResFailCode(c, common.StatusInvalidToken, common.MsgInvalidToken)
		//	return
		//}
		//
		////查询redis value是否匹配 key:userId value:token
		//if value, err := cache.GetString(userInfo["userId"]); err != nil {
		//	common.ResFailCode(c, common.StatusServerError, common.MsgInvalidToken)
		//	return
		//} else if token != value {
		//	//token不匹配, 期间帐号被他人登录,redis中的token刷新
		//	common.ResFailCode(c, common.StatusInvalidToken, common.MsgInvalidToken)
		//	return
		//}
		//
		////查询用户信息
		//userId := convert.ToInt(userInfo["userId"])
		//user, err := db.GetUserByUserId(userId)
		//if  err != nil {
		//	common.ResFailCode(c,common.StatusInvalidToken,common.MsgInvalidToken)
		//	return
		//}

		c.Set(UserKey, userInfo["userId"])
	}
}
