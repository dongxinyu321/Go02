package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func tongji(filename string, shuchu string) {
	ior, err := os.Open(filename)
	defer ior.Close()
	if err != nil {
	}
	reader := bufio.NewReader(ior)

	iow, err := os.OpenFile(shuchu, os.O_APPEND|os.O_CREATE, 0666)
	defer iow.Close()
	writer := bufio.NewWriter(iow)
	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err == io.EOF {
			break
		}
		//在寫入時追加換行符
		w, err := writer.Write(append([]byte("\n"), line...))
		fmt.Println("writer返回值", w)
		if err != nil && w == 0 {
		}
		writer.Flush()

	}

}

func main() {
	tongji("E:/test.txt", "E:/oooo.txt")
}
