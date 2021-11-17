package main

import (
	"fmt"
	"log_agent/etcd"
)

func main() {
	runApp()
}

func runApp()  {
	//初始化kafka

	// 初始化etcd

	//
	err := etcd.Watch("logs")
	if err != nil {
		fmt.Println("初始化监控失败", err)
		return
	}
}