package main

import (
	"fmt"
)

//接口的实现功能的组装

//类型断言 最佳实践

type Usb interface {
	Open()
	Close()
}

type Computer struct {
	Name string
}

func (c Computer) Open() {
	fmt.Println("电脑开机！！！")
}

func (c Computer) Close() {
	fmt.Println("电脑关机！！！")
}

type Phone struct {
	Name string
}

func (p Phone) Call() {
	fmt.Println("手机打电话！！！")
}

func (c Phone) Open() {
	fmt.Printf("%v手机开机！！！", c.Name)
}

func (c Phone) Close() {
	fmt.Printf("%v手机关机！！！", c.Name)
}

func Working(u Usb) {
	u.Open()
	u.Close()
	if p, info := u.(Phone); info == true {
		p.Call()
	}
}

func main() {
	var usbs [3]Usb
	usbs[0] = Phone{"小米"}
	usbs[1] = Phone{"华为"}
	usbs[2] = Computer{"联想"}

	for i := 0; i < 3; i++ {
		Working(usbs[i])
	}
}
