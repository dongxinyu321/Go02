package main

import (
	"fmt"
)

func sum(n1, n2 int) int {
	//当执行到defer时  会将defer后面的语句压入到独立的栈中 暂时不执行
	//当函数执行完毕后 再从defer栈 按照先入后出的方式出栈 执行 先进后出
	// 打开连接 锁后 可以直接defer 在defer后可以继续使用打开的资源
	//当函数执行完毕后系统会依次从defer栈中取出defer后语句关闭资源 好处能帮助处理关闭
	defer fmt.Println("okr n1=", n1)
	defer fmt.Println("olr n2=", n2)
	fmt.Println("okr3 =", n1+n2)
	return n1 + n2
}

func main() {
	fmt.Println(sum(10, 20))
}
