package main

import "fmt"

type person struct {
	name string
	age int
	gender string
	hobby []string
}


func main()  {
	// new  分配内存， 返回地址
	// make 分配内存， 返回引用
	var p = new(person)
	fmt.Println(p)

	p2 := person{"孙世军",18, "", nil}
	fmt.Println(p2)

	var  animal struct{
		name string
		foot string
	}
	animal.name = "abc"
	// 结构体是值类型

	p3 := newPerson("abc", 18, "男", []string{})
	p3.eating("土豆")

	fmt.Println(p3)
}


func changeGender(p *person) {
	(*p).gender = "男"
	p.gender = "女"
}
// 构造函数
func newPerson(name string, age int, gender string, hobby []string)*person{
	return &person{
		name: name,
		age: age,
		gender: gender,
		hobby: hobby,
	}
}


// 方法
func (p *person) eating(food string)  {
	p.name = "被修改了"
	fmt.Printf("%s在吃%s", p.name, food)
}