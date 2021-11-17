package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hpcloud/tail"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log_agent/agent"
	"log_agent/config"
	"time"
)

type LogType struct {
	Topic string `json:"topic"`
	Filename string `json:"filename"`
}

func Watch(key string) (err error)  {
	var cli *clientv3.Client
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{config.AppConfigObj.EtcdConfig.Address},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch key:q1mi change
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	// 一旦发现 配置变动， 就告诉log agent 去 读相应的日志， 并将督导的结果发送给kafka
	var logs = make([]LogType, 10)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			err = json.Unmarshal(ev.Kv.Value, &logs)
			if err != nil {
				fmt.Println("填入的json不合法：", err)
				return
			}
			if len(logs) <= 0 {
				fmt.Println("填入的json不能为空")
				return
			}
			fmt.Println(logs)
			for _, item := range logs {
				var tailInstance *tail.Tail
				go func(t *tail.Tail, l LogType) {
					t, err = agent.ReadFile(l.Filename)
					if err != nil {
						fmt.Println("读取日志失败")
						return
					}
					agent.Run(t, l.Topic)
				}(tailInstance, item)
			}
		}
	}
	return
}
