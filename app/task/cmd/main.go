package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"micro_play/app/task/repository/db"
	"micro_play/app/task/service"
	"micro_play/config"
	"micro_play/idl/pb"
)

func main() {

	config.InitConfig()
	db.InitDb()

	etcdReg := etcd.NewRegistry(registry.Addrs(fmt.Sprintf("%v:%v", config.GlobalConfig.EtcdConfig.Address, config.GlobalConfig.EtcdConfig.Port)))

	microService := micro.NewService(
		micro.Name(config.GlobalConfig.HostServiceName),
		micro.Address(config.GlobalConfig.HostServiceAddr),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())
	_ = microService.Run()
}
