package factory

import "fmt"

//工厂模式 构造方法

type student struct {
	Name string
	Age  int
}

//student 构造方法
func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func (s *student) SetAge(age int) {
	s.Age = age
}

func (s *student) GetName() string {
	return s.Name
}

type pipl struct {
	student
}

func NewPipl(name string, age int) *pipl {
	p := pipl{}
	p.SetName(name)
	p.SetAge(age)
	return &p
}

func (p *pipl) Kaoshi() {
	fmt.Println("小学生正在考试")
}
