package main

import (
	"fmt"
)

func doStuff(value []string) {
	fmt.Printf("value=%v\n", value)
	//value2 := value[:]
	var value2 []string
	//初始数组容量足够的情况下不会新建数组 会共用一个地址 造成原数组值得改变
	//copy(value,value2)
	value2 = append(value2, "b")
	fmt.Printf("value=%v, value2=%v\n", value, value2)
	value2[0] = "z"
	fmt.Printf("value=%v, value2=%v\n", value, value2)
}
func main() {
	//slice1 := []string{"a"} // 长度 1, 容量 1
	slice1 := make([]string, 1, 1)
	slice1[0] = "a"
	doStuff(slice1)
	// Output:
	// value=[a] -- ok
	// value=[a], value2=[a b] -- ok: value 未改变, value2 被更新
	// value=[a], value2=[z b] -- ok: value 未改变, value2 被更新
	slice10 := make([]string, 1, 10) // 长度 1, 容量 10
	slice10[0] = "a"
	doStuff(slice10)
	// Output:
	// value=[a] -- ok
	// value=[a], value2=[a b] -- ok: value 未改变, value2 被更新
	// value=[z], value2=[z b] -- WTF?!? value 改变了???
}
