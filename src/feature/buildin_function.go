package main

import "fmt"

func main() {
	// len(Type) 求长度

	// new(Type) 分配内存 主要分配给值(指针)类型
	var num *int = new(int)
	fmt.Printf("num的类型 = %T, num的值 = %v, num2的地址 = %v, num2指向的值 = %v", num, num, &num, *num)
}
