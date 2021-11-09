package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var notify = make(chan bool, 1)

func main() {
	wg.Add(1)
	go p()
	time.Sleep(time.Second * 5)
	notify <- true
	wg.Done()
}

func p() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("孙")
		time.Sleep(time.Millisecond * 1000)
		select {
		case <-notify:
			break LOOP
		default:
		}
	}
}
