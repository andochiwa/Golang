package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端开始连接")
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("client connect error", err)
		return
	}
	fmt.Println("client connect success", conn.RemoteAddr().String())
	defer func(dial net.Conn) {
		err := dial.Close()
		if err != nil {
			fmt.Println("client close error", err)
			return
		}
	}(conn)

	// 发送单行数据, os.Stdin表示标准输入
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	// 把line发送给服务器
	write, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("connect write error", err)
	}
	fmt.Printf("client send %d byte data, and exit\n", write)

}
