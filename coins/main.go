package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew",
		"Sarah",
		"Augustus",
		"Heidi",
		"Emilie",
		"Peter",
		"Giana",
		"Adriano",
		"Aaron",
		"Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

//自定义类型
type  myInt int
//  类型别名
type  yourInt = int
var a rune



func main()  {
   fmt.Println("剩下：", dispatchCoin())
	fmt.Println("分配情况：", distribution)
}

func dispatchCoin() int {
	for _, v:= range users {
		coins -= coinsPerPerson(v)
	}
	return coins
}

func coinsPerPerson(name string)  (coins int){
	if name != ""{
       for _,v := range name {
       	  if v=='e' || v =='E' {
       	  	coins += 1
		  }
		   if v=='i' || v =='I' {
			   coins += 2
		   }
		   if v=='o' || v =='O' {
			   coins += 3
		   }
		   if v=='u' || v =='U' {
			   coins += 4
		   }
	   }
	}
	distribution[name] = coins
	return
}
