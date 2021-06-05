package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"redis.demo/common/message"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("等待客户端发送数据")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read error =", err)
		return
	}
	// 根据buf[:4] 转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[:4])

	// 根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if err != nil || n != int(pkgLen) {
		return
	}
	// 把pkgLen反序列化成message.Message类型
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()
	// 读取客户端发送的信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("对方关闭了连接，服务正常退出")
				return
			}
			fmt.Println("readPkg err =", err)
			return
		}
		fmt.Println("mes =", mes)
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
