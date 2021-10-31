package main

import (
	"fmt"
	"time"
)

func main() {

	inchan := make(chan int, 10000000)
	start := time.Now().Unix()
	for i := 1; i <= 2000000; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i, "是素数")
			inchan <- i
		}
	}
	close(inchan)
	end := time.Now().Unix()
	for v := range inchan {
		fmt.Println(v)
	}
	fmt.Println(end - start)
}
