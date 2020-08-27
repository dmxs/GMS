package grpc

import (
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/proto/user"
	svc "GMS/srv/user/service"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"time"
)

func New(c *conf.Config) (svc service.Service){
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

	return grpc.NewService(
		//micro.Registry(registry),
		service.Name(conf.Conf.Micro.Name),
		service.Version(conf.Conf.Micro.Version),
		//micro.RegisterTTL(time.Second*15),
		service.RegisterInterval(time.Second*10),
		//micro.WrapHandler(opentracing.NewHandlerWrapper(openTrace.GlobalTracer())),
	)
}

func Start(grpc service.Service, service *svc.Service) {
	//handle
	if err := user.RegisterUserHandler(grpc.Server(), &server{ s:service } ); err != nil {
		panic(err)
	}

	if err := grpc.Run(); err != nil {
		logger.Error(err.Error())
		return
	}
}