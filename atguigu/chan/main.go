package main

import "fmt"

//协程
//管道
func main1() {

	var inChan chan int
	inChan = make(chan int, 3)

	//inChan <- 10
	num := 211
	inChan <- num
	fmt.Println(len(inChan), " 容量为:", cap(inChan))

	var num2 int
	num2 = <-inChan
	fmt.Println(num2)
	fmt.Println(len(inChan), " 容量为:", cap(inChan))

}
