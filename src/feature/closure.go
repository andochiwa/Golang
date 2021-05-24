package main

import "fmt"

// 闭包使用 累加器
func addUpper(n int) func(int) int {
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	// 因为go中函数为一等公民，本身可以作为一个数据类型，变量，模拟成类和对象等
	// 匿名函数和n形成一个整体构成闭包，所以相当于n已经初始化完毕
	// 也可以理解为闭包是一个类，匿名函数就是类的方法，而匿名外变量n就是类的字段
	f := addUpper(10) // 其实就相当于new了一个对象，这个对象就是addUpper函数本身
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(1)) // 14
}
