package main

import (
	"fmt"
	studentSMS "sms/str-sms"
)

type student struct {
	id   int
	name string
}

var students = make(map[int]*student)

func main() {
	/*var (
		input     int
		inputName string
		inputId   int
	)
	fmt.Println("欢迎使用学生管理系统：")
	for {
		fmt.Println("请输入数字选择：1.查看全部学生 2.添加学生（输入name 和 id）3.根据id 删除学生")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("输入不合法，请重新输入")
			continue
		}
		switch input {
		case 1:
			{
				fmt.Println("全部学生如下")
				getAllStudent()
			}
		case 2:
			{
				fmt.Println("请输入name 和 id")
				for {
					fmt.Scanf("%s %d", &inputName, &inputId)
					err := addStudent(inputName, inputId)
					if err != nil {
						fmt.Println(err, "请重新输入")
						continue
					}
					fmt.Println("添加成功")
					break
				}
			}
		case 3:
			{
				fmt.Println("请输入要删除的学生id")
				for  {
					fmt.Scanln(&inputId)
					err := rmStudentById(inputId)
					if err != nil {
						fmt.Println(err, "请重新输入")
						continue
					}
					fmt.Println("删除成功")
					break
				}
			}
		default:
			{
				fmt.Println("无效数字")
			}
		}
		fmt.Println("按回车返回主菜单")
		_,err1 := fmt.Scanln(&input)
		if err1 != nil {
			continue
		}
	}*/
    var input int
    system := studentSMS.NewSystem()
	fmt.Println("欢迎使用学生管理系统：")
	for {
		fmt.Println("请输入数字选择：1.查看全部学生 2.添加学生（输入name 和 id）3.根据id 删除学生")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("输入不合法，请重新输入")
			continue
		}
		switch input {
		case 1:
			system.ListStudents()
		case 2:
			system.AddStudent()
		case 3:
			system.RemoveStudentById()
		default:
			fmt.Println("无效数字")
		}
		fmt.Println("按回车返回主菜单")
		_,err1 := fmt.Scanln(&input)
		if err1 != nil {
			continue
		}
	}


}

func getAllStudent() {
	for _, v := range students {
		fmt.Println("name:", v.name, ",id:", v.id)
	}
}

func addStudent(name string, id int) error {
	if _, ok := students[id]; ok {
		return fmt.Errorf("添加失败：已存在id为%d的学生", id)
	}
	students[id] = &student{
		name: name,
		id:   id,
	}
	return nil
}

func rmStudentById(id int) error{
	if _, ok := students[id]; !ok {
		return fmt.Errorf("删除失败：不存在id为%d的学生", id)
	}
	delete(students, id)
	return nil
}
