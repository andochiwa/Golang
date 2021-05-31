package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() // 开启一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("main " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
