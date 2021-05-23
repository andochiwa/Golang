package main

import "fmt"

func main() {
	// go的类型转换需要显示转换
	var i int = 10
	// int ==> float32
	var f float32 = float32(i)
	fmt.Println("f=", f)

}
