package http

import (
	"GMS/pkg/convert"
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/service"
	svc "GMS/srv/user/service"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

var (
	svr *service.Service
)

func Init(c *conf.Config, service *svc.Service) {
	//web server
	svr = service
	webSvc := web.NewService(
		web.Name(conf.Conf.Micro.Name + ".http"),
		web.Address(":" + convert.ToString(conf.Conf.Http.Port)),
		web.Handler(NewRouter()),
		web.Metadata(map[string]string{"protocol" : "http"}),
	)

	logger.Info(conf.Conf.Micro.Name + ".http")

	if err := webSvc.Run(); err != nil {
		logger.Error(err.Error())
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("user/getByID", getByID)
	router.POST("user/getRole", getRole)

	return router
}