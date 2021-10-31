package main

//指针类型
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

//值传递
func main() {

}
