package Digui

import "fmt"

func Digui(n int) int {
	var j int = 1
	var m int = 1
	var o int
	for i := 1; i < n-1; i++ {
		o = m
		m = j + m
		j = o
	}
	return m
}

func Fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fbn(n-1) + Fbn(n-2)
	}
}

func F(n int) int {
	if n == 1 {
		return 3
	} else {
		return F(n-1) + 1
	}
}

func Sum(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T", n1)
	return n1 + n2
}
