package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// 获取并处理消息
func process(conn net.Conn) {
	defer conn.Close()
	// 读取客户端发送的信息
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("对方关闭了连接，服务正常退出")
				return
			}
			fmt.Println("readPkg err =", err)
			return
		}
		err = serverProssMessage(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMessage err =", err)
			return
		}
	}
}

// 根据消息种类分发函数
func serverProssMessage(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMessageType:
		// 处理登录
		err := ServerProcessLogin(conn, mes)
		if err != nil {
			return err
		}
	case message.RegisterMessageType:
		// todo 处理注册
	default:
		err = errors.New("消息类型不存在，无法处理")
		return
	}
	return
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
