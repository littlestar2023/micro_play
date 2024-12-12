package config

import (
	"github.com/littlestar2023/common_pkg/initialize"
)

var (
	GlobalConfig *ApplicationConfig
)

type ApplicationConfig struct {
	EtcdConfig EtcdConfig `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}

type EtcdConfig struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
}

func InitConfig() {

	GlobalConfig = &ApplicationConfig{}
	initialize.InitializeWithCustomize(GlobalConfig)
}
