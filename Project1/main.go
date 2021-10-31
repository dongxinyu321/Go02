package main

import (
	gn "Go02/Project1/gongneng"
	"fmt"
)

func main() {
	var b gn.Book
	var key string
	var ru string
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println()
		fmt.Println("1收支明细")
		fmt.Println("2登记收入")
		fmt.Println("3登记支出")
		fmt.Println("4退出")
		fmt.Println()
		fmt.Println("请选择(1-4)：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			b.Mingxi()
		case "2":
			b.Shouru()
		case "3":
			b.Zhichu()
		case "4":
			fmt.Println("是否確定退出y/n:")
			fmt.Scanln(&ru)
			switch ru {
			case "y":
				return
			case "Y":
				return
			case "n":
				break
			case "N":
				break
			}
		default:
			fmt.Println("输入错误请重新输入！！")
			continue
		}
	}
}
