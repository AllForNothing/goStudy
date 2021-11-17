package agent

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log_agent/kafka"
	"time"
)



func ReadFile(filename string) (tailInstance *tail.Tail, err error) {
	tailConfig := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	}
	tailInstance, err = tail.TailFile(filename, tailConfig)
	if err != nil {
		fmt.Println("初始化tail失败", err)
	}
	return
}




func Run(tail *tail.Tail, topic string)  {
	for  {
		line, ok := <-tail.Lines
		if !ok {
			fmt.Printf("读取不成功：%v", tail.Filename)
			time.Sleep(time.Second)
			continue
		}
		err := kafka.SendMsg(line.Text, topic)
		if err != nil {
			return
		}
	}
}
