package main

import (
	"GMS/pkg/logger"
	"GMS/srv/api/router"
	"github.com/micro/go-micro/v2/web"
)

func init() {
	logger.InitLog()
}

func main() {
	service := web.NewService(
		web.Name("gms.api.srv.service"),
		web.Address(":8080"),
		web.Handler(router.NewRouter()),
	)

	if err := service.Run(); err != nil {
		logger.Error("start failed : " + err.Error())
	}
}
