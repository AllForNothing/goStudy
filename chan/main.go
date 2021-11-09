package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jobChan := make(chan int64, 10)
	result := make(chan int64, 10)
	go random(jobChan)
	for i := 0;i<24;i++ {
		go pushToResult(jobChan, result, i)
	}
	/*for {
		ret, ok := <- result
		if !ok {
			break
		}
		fmt.Println("计算结果为：", ret)
	}*/
	for i:= range result{
		fmt.Println("计算结果为：", i)
	}
}

func random(jobChan chan<- int64)  {
	for {
		num := rand.Int63()
		jobChan <- num
		time.Sleep(time.Millisecond*500)
		fmt.Println("往jobChan里面发送随机数：", num)
	}
}

func pushToResult(jobChan <-chan int64, resultChan chan<- int64, jobId int)  {
	for {
		i,ok := <-jobChan
		if !ok {
			break
		}
		resultChan <- sum(i)
		fmt.Println("执行计算的JOB是：", jobId)
	}
}





func sum(num int64)  int64{
	count := int64(0)
	for num > 0 {
		count += num % 10
		num = num / 10
	}
	return count
}