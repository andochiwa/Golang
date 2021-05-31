package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 副协程
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("world")
		}
	}()
	// 主协程
	for i := 0; i < 10; i++ {
		// 因为go的协程是非抢占式的，让其变成强制式就是runtime.Gosched()方法
		//让出CPU时间片，重新等待安排任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
