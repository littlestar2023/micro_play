package config

import (
	"github.com/littlestar2023/common_pkg/initialize"
)

var (
	GlobalConfig *ApplicationConfig
)

type ApplicationConfig struct {
	EtcdConfig      EtcdConfig `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	HostServiceName string     `mapstructure:"host_service_name" json:"host_service_name" yaml:"host_service_name"`
	HostServiceAddr string     `mapstructure:"host_service_addr" json:"host_service_addr" yaml:"host_service_addr"`
}

type EtcdConfig struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
}

func InitConfig() {

	GlobalConfig = &ApplicationConfig{}
	initialize.InitializeWithCustomize(GlobalConfig)
}
