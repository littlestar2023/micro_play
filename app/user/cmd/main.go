package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"micro_play/app/user/service"
	"micro_play/config"
	"micro_play/idl/pb"
)

func main() {

	config.InitConfig()

	etcdReg := etcd.NewRegistry(registry.Addrs(fmt.Sprintf("%v:%v", config.GlobalConfig.EtcdConfig.Address, config.GlobalConfig.EtcdConfig.Port)))
	microService := micro.NewService(
		micro.Name(config.GlobalConfig.HostServiceName),
		micro.Address(config.GlobalConfig.HostServiceAddr),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	_ = microService.Run()
}