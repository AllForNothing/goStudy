package main

import (
	"fmt"
	"sync"
)

var wg  sync.WaitGroup


func hello(index int)  {
	defer wg.Done()
	fmt.Println("dfdfd",index)

}

func main() {
    c1 :=  make(chan int, 100)
	c2 :=  make(chan int, 100)
	wg.Add(1)
	go pushToChan(c1)
	wg.Add(1)
	go sq(c1, c2)
    wg.Wait()
	for i:= range c2 {
		fmt.Println(i)
	}
}

func pushToChan(cha chan int)  {
	defer wg.Done()
	for i:=0;i<100;i++ {
		cha <- i
	}
	close(cha)
}

func sq(cha1 chan int, cha2 chan int)  {
	defer wg.Done()
	for  {
		i,ok := <- cha1
		if !ok {
			break
		}
		cha2 <- i*i
	}
	close(cha2)
}