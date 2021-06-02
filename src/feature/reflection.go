package main

import (
	"fmt"
	"reflect"
)

type Student2 struct {
	Name  string
	Age   int
	Score int
}

func reflectTest(a interface{}) {
	// 通过反射获取type，kind，值
	// 获取Type
	rType := reflect.TypeOf(a)
	fmt.Println("rType =", rType)

	// 获取Value
	rValue := reflect.ValueOf(a)
	fmt.Println("rValue =", rValue)
}

func reflectTest2(a interface{}) {
	// 通过反射获取type，kind，值
	// 获取Type
	rType := reflect.TypeOf(a)
	fmt.Printf("rType = %v type = %T\n", rType, rType)

	// 获取Value
	rValue := reflect.ValueOf(a)
	fmt.Printf("rValue = %v, type = %T\n", rValue, rValue)
}

func main() {

	num := 100
	reflectTest(num)

	student := Student2{"andochiwa", 10, 50}
	reflectTest2(student)

}
