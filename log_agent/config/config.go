package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type KafkaConfig struct {
	Address string `ini:"address"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
}

// AppConfig 整个项目的配置
type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	EtcdConfig `ini:"etcd"`
}

var AppConfigObj = new(AppConfig)
func init() {
	err := ini.MapTo(AppConfigObj, "config/config.ini")
	fmt.Println("初始化对象为：", AppConfigObj)
	if err != nil {
		fmt.Println("初始化配置项出错：",err)
	}
}
