package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误信息：", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func main() {
	str := "字符串长度"
	//一个汉字占3个字节
	fmt.Println("str len", len(str))

	//如果字符串里有中文 先转化成
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符= %c\n", str2[i])
	}
	//字符串转整数： n ,err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转化错误：", err)
	} else {
		fmt.Println("打印转化数字: ", n)
	}

	//整数转字符串
	itoa := strconv.Itoa(12345)

	fmt.Printf("转字符串结果 %v 类型%T\n", itoa, itoa)
	//10进制 转2进制
	fmt.Printf("2进制= %v\n", strconv.FormatInt(123, 2))

	//查找字符串是否在在指定字符串中
	fmt.Printf("是否包含foo %v\n", strings.Contains("foodfoos", "foo"))

	//统计有几个不重复字符串
	fmt.Printf("num = %v\n", strings.Count("foodfoos", "o"))

	//将字符串左右两边的空格去掉
	space := strings.TrimSpace("! sda sda  dad !")
	fmt.Printf("space= %v\n", space)

	//将字符串左右两边指定的字符串去掉
	trim := strings.Trim(space, "!")
	fmt.Println(trim)

	test()

	for {
		fmt.Println("测试")
		time.Sleep(time.Second)
	}
}
