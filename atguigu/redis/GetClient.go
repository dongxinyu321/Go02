package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func init() {

}

func main() {
	var bt bytes.Buffer
	ss := make([]string, 100)
	valueList := make([]int64, 1000)
	options := redis.Options{Addr: "hadoop102:6379"}
	re := redis.NewClient(&options)
	pipeline := re.Pipeline()
	background := context.Background()

	for i := 0; i < 100; i++ {
		bt.WriteString("age")
		bt.WriteString(string(i))
		pipeline.Do(background, "Set", bt.String(), i)
		ss = append(ss, bt.String())
		bt.Truncate(0)
	}
	fmt.Println("set 完毕")

	for _, key := range ss {
		pipeline.Do(background, "Get", key)
		incr := pipeline.Incr(background, key)
		val := incr.Val()

		valueList = append(valueList, val)

	}
	fmt.Println("get 完毕")

	for _, va := range valueList {
		fmt.Println("开始遍历结果:", va)
	}
}
