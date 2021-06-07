package main

import (
	"fmt"
	"net"
	"redis.demo/common/utils"
)

// 和服务器保持通讯
func processServerMessage(conn net.Conn) {
	for {
		message, err := utils.ReadPkg(conn)
		if err != nil {
			fmt.Println("读取消息失败 =", err)
			return
		}
		fmt.Printf("收到消息 = %v\n", message.Data)
	}
}
