package studentSMS

import (
	"fmt"
)

type student struct {
	id   int
	name string
}

func NewSystem() *system {
	return &system{
		AllStudents: make(map[int]*student,50),
		startId: 202107080001,
	}
}

type system struct {
	AllStudents map[int] *student
	startId int
}

func (s *system) ListStudents() {
	for _,v := range s.AllStudents {
		fmt.Println(v.name, v.id)
	}
}


func (s *system) AddStudent() {
	var name string
	for  {
		for  {
			fmt.Println("请输入name")
			_, err := fmt.Scanln(&name)
			if err != nil {
				fmt.Println("输入的name不合法")
				continue
			}
			break
		}
		if _,ok := s.AllStudents[s.startId]; ok {
			fmt.Println("用户已存在")
			continue
		}
		s.AllStudents[s.startId] = &student{
			name: name,
			id: s.startId,
		}
		fmt.Println("添加成功")
		s.startId++
		break
	}
}

func (s *system) RemoveStudentById() {
	var id int
	for  {
		for  {
			fmt.Println("请输入id")
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("id")
				continue
			}
			break
		}
		if _,ok := s.AllStudents[id]; !ok {
			fmt.Println("用户不存在")
			continue
		}
		delete(s.AllStudents, id)
		fmt.Println("删除")
		break
	}
}