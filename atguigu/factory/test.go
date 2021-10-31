package factory

//单例工厂模式
type nv struct {
	nmae string
	age  int
	male string
}

func NewPerson(name string, age int, mael string) *nv {
	return &nv{
		nmae: name,
		age:  age,
		male: mael,
	}
}
