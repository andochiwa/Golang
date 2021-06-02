package main

import (
	"fmt"
	"reflect"
)

func reflectTest(a interface{}) {
	// 通过反射获取type，kind，值
	// 获取Type
	rType := reflect.TypeOf(a)
	fmt.Println("rType =", rType)

	// 获取Value
	rValue := reflect.ValueOf(a)
	fmt.Println("rValue =", rValue)
}

func main() {

	num := 100
	reflectTest(num)

}
