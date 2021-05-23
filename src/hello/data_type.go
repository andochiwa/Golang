package main

import (
	"fmt"
	"strconv"
)

func main() {
	// go的类型转换需要显示转换
	var i int = 10
	// int ==> float32
	var f float32 = float32(i)
	fmt.Println("f =", f)

	// 转string 第一种
	n1, n2, n3, n4 := 10, 12.34, true, 'h'
	str := fmt.Sprintf("%d %f %t %c", n1, n2, n3, n4)
	fmt.Println(str)

	// 第二种 使用strconv的函数
	str = strconv.FormatInt(int64(n1), 10)
	fmt.Println(str)
}
