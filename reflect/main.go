package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address string `ini:"address"`
	Port int `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func main() {
	var config MysqlConfig
	v := reflect.ValueOf(&config).Elem()
	fileObj, err := os.Open("./conf.ini")
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
		if strings.Index(line,"=") != -1 {
			keyValue := strings.Split(line, "=")
			key := keyValue[0]
			value := keyValue[1]
			fmt.Println(key, value)
			for i := 0; i < v.NumField(); i++ {
				field := v.Type().Field(i)
				if field.Tag.Get("ini") == key {
					v.FieldByName(field.Name).Set(reflect.ValueOf(convertStringByKind(value, field.Type)))
				}
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("读文件出错：%v\n", err)
			return
		}
	}

   fmt.Println(config)
}


func convertStringByKind(s string, t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Int:
		i, err := strconv.Atoi("8888")
		if err != nil {
			panic(err)
		}
		return i
	case reflect.String:
		return s
	}
	return nil
}

