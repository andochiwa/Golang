package main

import (
	"fmt"
	"math"
)

// 判断是否为素数
func checkPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		// 取不到就退出
		if !ok {
			break
		}
		if num <= 2 {
			primeChan <- num
		} else {
			flag := true
			sqrt := int(math.Sqrt(float64(num)))
			for i := 2; i <= sqrt; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				primeChan <- num
			}
		}
	}
	fmt.Println("其中一个协程退出了")
	// 这里还不能关闭管道，因为可能还有其他协程在运行
	exitChan <- true
}

// 判断1-8000数中的素数有哪些
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 3000) // 放入结果
	exitChan := make(chan bool, 4)    // 标识退出的管道

	// 开启协程，往管道中放入1-8000数
	go func(intChan chan int) {
		for i := 1; i <= 8000; i++ {
			intChan <- i
		}
		close(intChan)
	}(intChan)

	// 开启四个协程，判断是否为素数，如果是就放入到primeChan管道中
	for i := 0; i < 4; i++ {
		go checkPrime(intChan, primeChan, exitChan)
	}

	// 从退出标识管道中取值，如果取到了四个，说明所有协程计算已完成
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()

	for num := range primeChan {
		fmt.Printf("素数 = %v\n", num)
	}
}
