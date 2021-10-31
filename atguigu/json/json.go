package main

import (
	"encoding/json"
	"fmt"
)

type preson struct {
	Name string
	Age  int
	male string
}

func NewPerson(name string, age int, xingbie string) *preson {
	return &preson{
		Name: name,
		Age:  age,
		male: xingbie,
	}
}

//字符串序列化成json字符串
func main() {
	p := NewPerson("niu", 18, "男")
	p1, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("序列化错误为%v")
	}
	fmt.Println(string(p1))
}
