package main

import (
	"fmt"
)

func main() {
	go func() {
		defer fmt.Println("first defer")
		func() {
			defer fmt.Println("second defer")
			// 结束当前协程方法
			//runtime.Goexit()
			defer fmt.Println("third defer")
		}()
		fmt.Println("Function")
	}()
	for {
	}
}
