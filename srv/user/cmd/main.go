package main

import (
	"GMS/pkg/common"
	"GMS/pkg/logger"
	"GMS/srv/user/conf"
	"GMS/srv/user/dao"
	"GMS/srv/user/router"
	"github.com/micro/go-micro/v2/web"
	"os"
)

var (
	cfg *conf.Config
)

func init() {
	logger.InitLog()

	appPath ,_:= common.GetCurrentPath()
	logger.Info(appPath)
	if err := conf.LoadGlobalConfig(appPath + string(os.PathSeparator) + "config.toml"); err != nil {
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
	//grpcService := grpc.NewService(
	//	//micro.Registry(registry),
	//	service.Name(cfg.Micro.Name),
	//	service.Version(cfg.Micro.Version),
	//	//micro.RegisterTTL(time.Second*15),
	//	service.RegisterInterval(time.Second*10),
	//	//micro.WrapHandler(opentracing.NewHandlerWrapper(openTrace.GlobalTracer())),
	//)
	//
	//if err := user.RegisterUserHandler(grpcService.Server(), srv.New(dao.GetDao())); err != nil {
	//	logger.Error(err.Error())
	//	return
	//}

	//grpcService.Run()

	//if err := grpcService.Run(); err != nil {
	//	logger.Error(err.Error())
	//	return
	//}

	webService := web.NewService(
		web.Name(cfg.Micro.Name),
		web.Address(":8081"),
		web.Handler(router.NewRouter()),
	)

	if err := webService.Run(); err != nil {
		logger.Error(err.Error())
		return
	}
}
