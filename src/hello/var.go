package main

import "fmt"

// 声明多个全局变量
var (
	m1 = 10
	m2 = 10.34
	m3 = "hello"
)

func main() {
	// 声明变量
	var i int
	i = 10
	// 类型推导
	var num = 10.11
	fmt.Println("i =", i, "num =", num)

	// 省略var
	name := "hello"
	fmt.Println(name)

	// 声明多个变量
	var n1, n2, n3 int
	var n4, n5, n6 = 5.7, "ok", -1
	n7, n8, n9 := 10, 11, 12
	fmt.Println(n1, n2, n3, n4, n5, n6, n7, n8, n9)
	fmt.Println(m1, m2, m3)
}
