package main

import (
	"fmt"
	"net"
)

// Process 处理客户端传来的数据
func Process(conn net.Conn) {
	// 循环接受客户端发送的数据
	// 首先延迟关闭连接
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close client error", err)
		}
		fmt.Println("close client")
	}(conn)

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		fmt.Printf("等待客户端 %s 发送数据\n", conn.RemoteAddr().String())
		// n代表读取到的数据的大小
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client has exited")
			return
		}
		// 把信息显示在终端
		fmt.Printf("Server receive = %v\n", string(buf[:n]))
	}
}

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
		fmt.Printf("Accept success con = %v, ip = %v\n", accept, accept.RemoteAddr().String())

		// 启动一个协程处理信息
		go Process(accept)

	}
	//fmt.Println("listen success", listen)
}
