package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

////文件不存在方法
//var filename = "E:/test1.txt"
//_, err := os.Stat(filename)
//if err == nil{
//	fmt.Println(filename,"文件存在")
//}else if os.IsNotExist(err){
//	fmt.Println("文件不存在",err)
//}else{
//	fmt.Printf("其他错误",err)
//}

//copy 文件的几种方式
func CopyFile(rufile string, chufile string) (writer int64, err error) {

	read, err := os.Open(rufile)
	defer read.Close()
	if err == nil {
		fmt.Println("输入文件名没有错误")
	}
	reader := bufio.NewReader(read)

	write, err := os.OpenFile(chufile, os.O_WRONLY|os.O_CREATE, 0666)
	defer write.Close()
	if err != nil {
	}
	newWriter := bufio.NewWriter(write)

	written, err := io.Copy(newWriter, reader)
	return written, err

}

func CopyFile1(rufile string, chufile string) (err error) {

	file, err := ioutil.ReadFile(rufile)
	if err != nil {
	}
	err = ioutil.WriteFile(chufile, file, fs.ModeType)
	return err

}

func main() {
	//拷贝文件
	fuFile := "E:/截屏/tupian.jpg"
	chuFile := "E:/tu3.jpg"
	CopyFile1(fuFile, chuFile)

}
