package main

import (
	"fmt"
	"reflect"
)

func reflectSetGenetic(a interface{}) {
	// 获取到reflect.Value
	rVal := reflect.ValueOf(a)
	fmt.Printf("rVal kind = %v\n", rVal.Kind())
	// Elem方法会返回指针指向的具体的值，不然的话无法直接修改指针
	rVal.Elem().SetInt(20)
}

func main() {
	num := 10
	fmt.Println("修改前num =", num)
	reflectSetGenetic(&num)
	fmt.Println("修改后num =", num)

}
