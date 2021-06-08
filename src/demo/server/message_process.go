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
	var user User
	// 读取客户端发送的信息
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("对方关闭了连接，服务正常退出")
				userManager.DeleteOnlineUser(user)
				NotifyUsers(user.UserId, user.UserName, message.UserOffline)
				return
			}
			userManager.DeleteOnlineUser(user)
			NotifyUsers(user.UserId, user.UserName, message.UserOffline)
			fmt.Println("readPkg err =", err)
			return
		}
		user, err = serverProssMessage(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessMessage err =", err)
		}
	}
}

// 根据消息种类分发函数
func serverProssMessage(conn net.Conn, mes *message.Message) (user User, err error) {
	switch mes.Type {
	case message.LoginMessageType:
		// 处理登录
		user, err = ServerProcessLogin(conn, mes)
		if err != nil {
			return
		}
	case message.RegisterMessageType:
		// 处理注册
		err = ServerProcessRegister(conn, mes)
		if err != nil {
			return
		}
	case message.SmsMessageType:
		// 处理用户发送消息
		err = SendMessageToUsers(mes)
		if err != nil {
			return
		}
	default:
		err = errors.New("消息类型不存在，无法处理")
		return
	}
	return
}
