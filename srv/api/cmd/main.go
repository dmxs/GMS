package main

import (
	"GMS/pkg/logger"
	"GMS/srv/api/conf"
	"GMS/srv/api/http"
)

func main() {
	//log
	logger.InitLog()

	//config
	if err := conf.InitConfig(); err != nil {
		panic("config load failed " + err.Error())
	}

	//开启http服务
	http.Init(conf.Conf)
}
