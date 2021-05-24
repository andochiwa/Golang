package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然为数据类型，那也可以作为形参的参数
func getSumFun(sum func(int, int) int, n1 int, n2 int) int {
	return sum(n1, n2)
}

// 可以自定义类型为函数
type myFunc func(int, int) int

func getCustomizeSum(sum myFunc, n1 int, n2 int) int {
	return sum(n1, n2)
}

// 函数也支持对返回值命名
func getSubAndPro(n1 int, n2 int) (sub int, pro int) {
	sub = n1 - n2
	pro = n1 * n2
	return
}

func main() {
	// 函数也是一种数据类型，变量可以赋值为函数
	a := getSum
	fmt.Printf("a类型为%T getSum类型为%T\n", a, getSum)
	// getSum(10, 40)
	sum := a(10, 40)
	fmt.Printf("函数作为变量 sum = %d\n", sum)
	// 通过传入函数类型形参
	sum = getSumFun(a, 10, 50)
	fmt.Println("函数作为形参 sum =", sum)
	// 通过自定义函数形参
	sum = getCustomizeSum(a, 20, 70)
	fmt.Println("自定义函数类型 sum =", sum)
	// 函数对返回值命名
	sub, pro := getSubAndPro(30, 10)
	fmt.Println("函数对返回值命名 sub =", sub, "pro =", pro)
}
