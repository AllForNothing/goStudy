package main

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	name string
	age int
	isOk bool
)

const (
	n1 = iota
	n2
	n3
)

func main()  {
	s1 := `fsdf`
	length := len(s1)
	fmt.Printf("长度：%d", length)

	fmt.Println(strings.Split(s1, ""))


	intType := 10
	flotype := 3.14
	boolType := true
	stringType := "孙世军"

	fmt.Printf("类型：%T， 值：%d\n", intType, intType)
	fmt.Printf("类型：%T， 值：%f\n", flotype, flotype)
	fmt.Printf("类型：%T， 值：%t\n", boolType, boolType)
	fmt.Printf("类型：%T， 值：%s\n", stringType, stringType)

    s2 := "hello沙河小王子"
    count := 0
    for _, s := range s2 {
        if unicode.Is(unicode.Han, s) {
        	count ++
		}
	}

	fmt.Printf("汉子数量为：%d", count)

    // 数组定义   数组是值类型


    var srr [5]float64
    fmt.Println(srr)

    array := [...]int{1, 3, 5, 7, 8}

    var sum int

    for _, item := range array{
    	sum += item
	}
    fmt.Println(sum)
	begin := 0
	end := len(array) -1
    for ; begin < end;{
    	if array[begin] + array[end] <= 8 {
    		if array[begin] + array[end] == 8 {
				fmt.Printf("(%d,%d)\n", begin, end)
			}
			begin ++
		}
		if array[begin] + array[end] > 8 {
			end --
		}
	}
}