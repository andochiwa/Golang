package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听")
	// 表示使用tcp协议监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	// 延时关闭
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			return
		}
	}(listen)

	// 等待客户端来连接
	for {
		fmt.Println("等待客户端连接")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			continue
		}
		fmt.Println("Accept success", accept)

	}
	//fmt.Println("listen success", listen)
}
