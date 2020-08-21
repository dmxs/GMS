package main

import (
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/proto"
	"GMS/srv/user/router"
	srv "GMS/srv/user/service"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	//log
	logger.InitLog()

	//config
	if err := conf.InitConfig(); err != nil {
		panic("config load failed")
	}

	//修改Etcd地址函数
	//addrEtcd := func(opts *registry.Options) {
	//	opts.Addrs = cfg.Etcd.Addrs
	//}
	//
	//registry := etcdv3.NewRegistry(addrEtcd)
	//
	////设置opentrace
	//t, io, err := tracer.NewTracer(cfg.Jaeger.ServiceName, cfg.Jaeger.URL)
	//if err != nil {
	//	logger.Fatalln(err)
	//}
	//defer io.Close()
	//openTrace.SetGlobalTracer(t)

	//rpc server
	grpcSvc := grpc.NewService(
		//micro.Registry(registry),
		service.Name(conf.Conf.Micro.Name),
		service.Version(conf.Conf.Micro.Version),
		//micro.RegisterTTL(time.Second*15),
		service.RegisterInterval(time.Second*10),
		//micro.WrapHandler(opentracing.NewHandlerWrapper(openTrace.GlobalTracer())),
	)

	//handle
	if err := user.RegisterUserHandler(grpcSvc.Server(), srv.New(conf.Conf)); err != nil {
		logger.Error(err.Error())
		return
	}

	//web server
	webSvc := web.NewService(
		web.Name(conf.Conf.Micro.Name + ".http"),
		web.Address(":8081"),
		web.Handler(router.NewRouter()),
	)

	go grpcSvc.Run()

	webSvc.Run()
}
