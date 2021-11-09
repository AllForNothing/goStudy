package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var  count int
var wg sync.WaitGroup
var so sync.Once

var m = sync.Map{}
func main() {
	wg.Add(21)

	for i :=0;i<21;i++ {
		go func(n int) {
			key := strconv.Itoa(n)
            m.Store(key,n)
			fmt.Println(m.Load(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func add() {
	for i :=0;i<100;i++ {
		count +=1
		fmt.Println(count)
		time.Sleep(time.Millisecond *50)
	}
	wg.Done()
}