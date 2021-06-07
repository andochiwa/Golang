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
	var userId int
	var userName string
	// 读取客户端发送的信息
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("对方关闭了连接，服务正常退出")
				userManager.DeleteOnlineUser(userId)
				NotifyUsers(userId, userName, message.UserOffline)
				return
			}
			fmt.Println("readPkg err =", err)
			return
		}
		userId, userName, err = serverProssMessage(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMessage err =", err)
			return
		}
	}
}

// 根据消息种类分发函数
func serverProssMessage(conn net.Conn, mes *message.Message) (userId int, name string, err error) {
	switch mes.Type {
	case message.LoginMessageType:
		// 处理登录
		userId, name, err = ServerProcessLogin(conn, mes)
		if err != nil {
			return
		}
	case message.RegisterMessageType:
		err := ServerProcessRegister(conn, mes)
		if err != nil {
			return
		}
	default:
		err = errors.New("消息类型不存在，无法处理")
		return
	}
	return
}
