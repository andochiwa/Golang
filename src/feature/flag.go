package main

import (
	"flag"
	"fmt"
)

func main() {
	// 用flag解析命令行参数
	user := ""
	pwd := ""
	host := ""
	port := 0

	// 1 接受的变量
	// 2 -u后的指定参数
	// 3 如果没有指定参数的默认值
	// 4 说明
	flag.StringVar(&user, "u", "", "用户名为空")
	flag.StringVar(&pwd, "pwd", "", "密码为空")
	flag.StringVar(&host, "h", "localhost", "ip地址")
	flag.IntVar(&port, "port", 3306, "端口号")
	// 必须调用该方法进行转换
	flag.Parse()

	fmt.Printf("username = %v, password = %v, host = %v, port = %v",
		user, pwd, host, port)
}
