package main

import "fmt"

// 全局变量优先初始化
var a = initVar()

func initVar() int {
	fmt.Println("init variable")
	return 0
}

// 每个源文件都可以包含一个init函数，会在main函数执行前被go运行框架调用
func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("main function")

	// 匿名函数，求两个数的值
	sum := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("匿名函数 sum =", sum)
}
