package main

import (
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/server/grpc"
	"GMS/srv/user/server/http"
	"GMS/srv/user/service"
)

func main() {
	//log
	logger.InitLog()

	//config
	if err := conf.InitConfig(); err != nil {
		panic("config load failed")
	}

	//rpc server
	grpcSvr := grpc.New(conf.Conf)

	//service
	svc := service.New(conf.Conf, grpcSvr.Client())

	//开启grpc 服务
	go grpc.Start(grpcSvr, svc)

	//开启http 服务
	http.Init(conf.Conf, svc)
}
