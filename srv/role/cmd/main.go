package main

import (
	"GMS/pkg/logger"
	"GMS/srv/role/conf"
	"GMS/srv/role/server/grpc"
	"GMS/srv/role/service"
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
	grpc.Start(grpcSvr, svc)
}
