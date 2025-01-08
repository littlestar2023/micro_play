package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"micro_play/app/gateway/router"
	"micro_play/app/gateway/rpc"
	"micro_play/config"
	"time"
)

func main() {

	config.InitConfig()

	rpc.InitRPC()

	etcdReg := etcd.NewRegistry(registry.Addrs(fmt.Sprintf("%v:%v", config.GlobalConfig.EtcdConfig.Address, config.GlobalConfig.EtcdConfig.Port)))

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name(config.GlobalConfig.HostServiceName),
		web.Address(config.GlobalConfig.HostServiceAddr),
		// 将服务调用实例使用gin处理
		web.Handler(router.InitRouters()),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
