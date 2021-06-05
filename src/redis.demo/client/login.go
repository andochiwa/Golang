package main

import (
	"encoding/json"
	"net"
	"redis.demo/common/message"
)

// 登录函数
func login(id int, password string) error {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		return err
	}

	// 发送消息给服务器
	msg := message.Message{}
	msg.Type = message.LoginMessageType
	// 创建一个LoginMessage结构体
	loginMsg := message.LoginMessage{}
	loginMsg.UserId = id
	loginMsg.UserPwd = password

	// 将loginMsg序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		return err
	}
	// 将序列化好的消息转给msg
	msg.Data = string(data)

}
