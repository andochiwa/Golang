package main

import "fmt"

func main() {
	fmt.Println("hello\nworld")
	//sum, sub := getSumAndSub(1, 2)
}

// 多个返回值
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// 指针
func testPtr(num *int) {
	*num = 20
}
