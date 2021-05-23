package main

import "fmt"

func main() {
	// 获取终端输入数据
	var name string
	var age int
	fmt.Println("请输入姓名 年龄 按回车结束")
	fmt.Scanln(&name, &age)
	fmt.Println("姓名 =", name, "年龄 =", age)
}
