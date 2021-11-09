package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	readByIOUtil()
}


//根据字节来读
func readByByte()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
		return
	}
	defer func() {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("关闭文件出错：%v\n", err)
		}
	}()
	for  {
		tmp := make([]byte,128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("读文件出错：%v\n", err)
			return
		}
		fmt.Print(string(tmp))
		// 说明读完了
		if n < 128 {
			return
		}
	}
}

// 使用bufferIO来读
func readByBuffer()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
		return
	}
	defer func() {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("关闭文件出错：%v\n", err)
		}
	}()
	reader := bufio.NewReader(fileObj)
	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("读文件出错：%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

//使用ioUtil来读
func readByIOUtil()  {
	file, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
		return
	}
	fmt.Println(string(file))
}

// 读取文件建议使用bufio  的 ReadLine  方式