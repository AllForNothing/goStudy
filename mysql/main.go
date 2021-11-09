package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sdn := "root:123456@tcp(10.202.250.196:3306)/test"

	db, err := sql.Open("mysql", sdn)
	if err != nil {
		fmt.Printf("sdn 格式不正确：%v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("sdn 格式不正确：%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
}
