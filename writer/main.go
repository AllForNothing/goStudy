package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeByUtil()
}

// 通过openFile  写文件
func writeByOS()  {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("关闭文件出错：%v\n", err)
		}
	}()
	_, err1 := file.WriteString("我是第一行\n")
	if err1 != nil {
		fmt.Printf("写入文件出错：%v\n", err1)
		return
	}
}

// 通过buffer 写
func writeByBuffer()  {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("关闭文件出错：%v\n", err)
		}
	}()
	writer := bufio.NewWriter(file)
	_, err1 := writer.WriteString("我是第二行\n")
	if err1 != nil {
		return
	}
	err2 := writer.Flush()
	if err2 != nil {
		return
	}
}

//通过util
func writeByUtil()  {
	err := ioutil.WriteFile("./xx.txt", []byte("我是第三行\n"), 0644)
	if err != nil {
		return 
	}
}