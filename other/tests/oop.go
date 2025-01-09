package main

import "fmt"

// 类：结构体
type Student struct {
	id uint
	name string
	male bool
	score float64
}

// 构造函数

func NewStudent(id uint, name string, score float64) *Student  {
	return &Student{
		id:    id,
		name:  name,
		score: score,
	}
}

// 成员方法

func (s Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}


func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}

func main()  {
	student := NewStudent(1, "张三", 100)
	fmt.Println(student) // &{1 张三 false 100}

	fmt.Println("Name: ", student.GetName()) // Name:  张三

	student.SetName("李四")
	fmt.Println("Name: ", student.GetName())

}

