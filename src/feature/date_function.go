package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now() // now type = time.Time
	fmt.Printf("now = %v, now type = %T\n", now, now)
	// 通过now可以获取年月日时分秒等
	fmt.Println("年 =", now.Year())
	fmt.Println("月 =", now.Month())
	fmt.Println("日 =", now.Day())

	// 格式化日期 至于为什么要这么写，我tm也不知道
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}
