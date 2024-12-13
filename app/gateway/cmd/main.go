package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"micro_play/config"
)

func main() {

	config.InitConfig()

	etcdReg := etcd.NewRegistry(registry.Addrs(fmt.Sprintf("%v:%v", config.GlobalConfig.EtcdConfig.Address, config.GlobalConfig.EtcdConfig.Port)))

}
