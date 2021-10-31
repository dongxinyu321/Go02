package main

import (
	"fmt"
)

//闭包
//全局匿名函数
var (
	Fun1 = func(n1, n2 int) int { return n1 + n2 }
)

//init在main方法前被执行
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	fmt.Printf(string(Fun1(10, 20)))
	//使用前面的代码
}
