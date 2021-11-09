package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "how do you do"
	arr := strings.Split(s, " ")
	map1 := make(map[string]int, len(s))

	for _, word :=range arr {
		_, ok := map1[word]
		if ok {
			map1[word]++
		} else {
			map1[word] = 1
		}
	}

	for k, v := range map1 {
		fmt.Printf("%s: %d\n", k, v)
	}

	s2 := "lo上海自来水来自海上ol"
	flag := true
	s2Arr := strings.Split(s2 ,"")
	for i :=0;i <= len(s2Arr)/2;i++ {
		if s2Arr[i] != s2Arr[len(s2Arr)-1-i] {
			flag = false
		}
	}
	fmt.Println(flag)

}

func b()  {

}

func a() {
	b()
}


