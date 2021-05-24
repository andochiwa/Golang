package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然为数据类型，那也可以作为形参的参数
func getSumFun(sum func(int, int) int, n1 int, n2 int) int {
	return sum(n1, n2)
}

func main() {
	// 函数也是一种数据类型，变量可以赋值为函数
	a := getSum
	fmt.Printf("a类型为%T getSum类型为%T\n", a, getSum)
	// getSum(10, 40)
	sum := a(10, 40)
	fmt.Printf("sum = %d\n", sum)
	// 通过传入函数类型形参
	sum = getSumFun(a, 10, 50)
	fmt.Println("sum =", sum)
}
