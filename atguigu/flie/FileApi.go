package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main1() {

	//打开文件
	//open, err := os.Open("E:/test.txt")
	//if err != nil{
	//	fmt.Println("打开文件出现异常",err)
	//}
	//defer open.Close()
	//reader := bufio.NewReader(open)
	//for {
	//	//_, str, err := reader.ReadLine()
	//	str, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		//fmt.Println("输出错误:",err)
	//		break
	//	}
	//	fmt.Println(str)
	//}
	//直接读取文件
	//str, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(str))
	//追加写入 加 读取
	filename := "E:/test.txt"
	shuchuname := "E:/test1.txt"
	file, err := os.OpenFile(filename, os.O_APPEND, 0644)
	defer file.Close()
	bytes := []byte("\n我要吃饭")
	write, err := file.Write(bytes)
	if err != nil && write < len(bytes) {
	}

	readFile, err := ioutil.ReadFile(filename)
	if err != nil {
		println("错误", err)
	}
	fmt.Println(string(readFile))

	err = ioutil.WriteFile(shuchuname, readFile, 0644)
	if err != nil {
		fmt.Println("输出有误", err)
	}

	i, err := ioutil.ReadFile(shuchuname)

	if err != nil {
		fmt.Println("读取文件", shuchuname, "有误")
	}
	fmt.Println("输出", shuchuname, "文件内容为:")
	fmt.Println(string(i))
}
