package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	p := person{
		"孙",
		18,
	}
	marshal, err := json.Marshal(p)
	if err != nil {
		return
	}
	s := string(marshal)
	fmt.Println(s)
	jsonStr := `{"name":"孙世军","age":18}`
	var p1 person
	err1 := json.Unmarshal([]byte(jsonStr), &p1)
	if err1 != nil {
		return
	}
	fmt.Println(p1)
}
