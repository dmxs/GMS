package http

import (
	"GMS/srv/user/service"
	"github.com/gin-gonic/gin"
)

var (
	svr *service.Service
)

func NewRouter(s *service.Service) *gin.Engine {
	svr = s

	router := gin.Default()

	router.POST("user/getByID", getByID)
	router.POST("user/getRole", getRole)

	return router
}