package main

import (
	"fmt"
	"net"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// 和服务器保持通讯
func processServerMessage(conn net.Conn) {
	for {
		msg, err := utils.ReadPkg(conn)
		if err != nil {
			fmt.Println("读取消息失败 =", err)
			return
		}
		switch msg.Type {
		case message.NotifyUserStatusType:
			// 有人上线或离线了
		default:
			fmt.Println("收到消息 =", msg.Data)
		}
	}
}
