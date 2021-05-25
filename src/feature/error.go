package main

import (
	"errors"
	"fmt"
)

func test() int {
	// 使用defer + recover处理异常
	defer func() {
		err := recover() // 内置函数 可以捕获到异常
		if err != nil {  // 说明捕获到异常
			fmt.Println("err =", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	return res
}

// 创建自定义错误
func customizeError() (err error) {
	return errors.New("自定义错误")
}

func main() {
	// 测试错误
	test()

	// 创建自定义错误
	err := customizeError()
	if err != nil {
		fmt.Println(err)
	}

	println("可以执行到这里")
}
