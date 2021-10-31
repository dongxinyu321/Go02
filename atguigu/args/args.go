package main

import (
	"fmt"
	"os"
)

//獲取命令行參數

func main() {
	//fmt.Println(len(os.Args))

	for i, v := range os.Args {
		fmt.Println(i, "    ", v)
	}
	//flag包解析命令行參數
}
