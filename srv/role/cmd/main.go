package main

import (
	"GMS/pkg/logger"
	"GMS/srv/role/conf"
	"GMS/srv/role/proto/role"
	srv "GMS/srv/role/service"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"time"
)

func main() {
	//log
	logger.InitLog()

	//config
	if err := conf.InitConfig(); err != nil {
		panic("config load failed")
	}

	//rpc server
	svc := grpc.NewService(
		//micro.Registry(registry),
		service.Name(conf.Conf.Micro.Name),
		service.Version(conf.Conf.Micro.Version),
		//micro.RegisterTTL(time.Second*15),
		service.RegisterInterval(time.Second*10),
		//micro.WrapHandler(opentracing.NewHandlerWrapper(openTrace.GlobalTracer())),
	)
	//handle
	if err := role.RegisterRoleHandler(svc.Server(), srv.New(conf.Conf)); err != nil {
		logger.Error(err.Error())
		return
	}

	svc.Run()
}
