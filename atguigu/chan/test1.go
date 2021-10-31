package main

import (
	"fmt"
	"time"
)

func putNum(intchan chan int, c chan int) {
	//开启协程向 intChan 放任1-800000
	for i := 0; i <= 2000000; i++ {
		intchan <- i
	}
	close(intchan)

}

func getNum(intchan chan int, outchan chan int, c chan int) {

	for {

		m, end := <-intchan

		flag := true
		if !end {
			break
		}
		for i := 2; i < m; i++ {
			if m%i == 0 {
				flag = false
				break
			}
		}
		if flag {

			outchan <- m
		}

	}

	fmt.Println("get完成")
	c <- 1
}

func main() {
	var num int
	intChan := make(chan int, 1000)
	outchan := make(chan int, 2000)
	exitChan := make(chan int, 100)
	num = 16

	start := time.Now().Unix()
	//fmt.Printf("%T",start)
	var end int64
	go putNum(intChan, exitChan)

	for i := 0; i < num; i++ {
		go getNum(intChan, outchan, exitChan)
	}

	go func() {
		for i := 0; i < num; i++ {
			<-exitChan
		}

		end = time.Now().Unix()

		close(outchan)
		fmt.Println("关闭结果协程")
		close(exitChan)

	}()

	for v := range outchan {
		fmt.Println("遍历", v)
	}
	fmt.Println("输出end", end)
	fmt.Println("输出start", start)
	fmt.Println("输出结果", end-start)
}
