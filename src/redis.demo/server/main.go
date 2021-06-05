package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// 读取客户端发送的信息
	for {
		buf := make([]byte, 4096)
		fmt.Println("等待客户端发送数据")
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read error =", err)
			return
		}
		fmt.Println("读取到的数据 =", buf[:4])
	}
}

func main() {

	fmt.Println("服务器在9876端口开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:9876")
	if err != nil {
		fmt.Println("net listen err =", err)
		return
	}

	// 监听成功就等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err =", err)
		}

		// 连接成功后启动协程和客户端保持通讯
		go process(conn)

	}

}
