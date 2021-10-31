package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

//决定方法决定你使用什么标准进行排序

//sort 进行排序

func (hs HeroSlice) Less(i, j int) bool {
	return hs[j].Age >= hs[i].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hero := Hero{}
	hero = hs[i]
	hs[i] = hs[j]
	hs[j] = hero
}

func main() {
	var hsl HeroSlice
	for i := 0; i < 10; i++ {
		hsl = append(hsl, Hero{Name: fmt.Sprintf("英雄%d", rand.Intn(100)), Age: rand.Intn(100)})
	}
	sort.Sort(hsl)
	fmt.Println(hsl)
}
