package main

import (
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/dao"
	"GMS/srv/user/proto"
	"GMS/srv/user/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"time"
)

var (
	cfg *conf.Config
)

func init() {
	logger.InitLog()

	if err := conf.LoadGlobalConfig("config.toml"); err != nil {
		logger.Error(err.Error())
		return
	}

	cfg = conf.GetGlobalConfig()

	dao.Init()
}

func main() {
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

	srv := grpc.NewService(
		//micro.Registry(registry),
		micro.Name(cfg.Micro.Name),
		micro.Version(cfg.Micro.Version),
		//micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		//micro.WrapHandler(opentracing.NewHandlerWrapper(openTrace.GlobalTracer())),
	)

	if err := user.RegisterUserHandler(srv.Server(), service.New(dao.GetDao())); err != nil {
		logger.Error(err.Error())
		return
	}

	if err := srv.Run(); err != nil {
		logger.Error(err.Error())
		return
	}
}
