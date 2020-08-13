package middleware

import (
	"github.com/gin-gonic/gin"
)

var (
	cors = map[string]bool{"*": true}
)

func CrossOriginMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		// 跨域处理
		if origin := c.GetHeader("Origin"); cors[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
		} else if len(origin) > 0 && cors["*"] {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
}
