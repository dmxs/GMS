package router

import (
	"GMS/pkg/common"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("user/get", func(c *gin.Context) {
		common.ResSuccessMsg(c)
		return
	})

	return router
}