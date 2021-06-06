package main

import (
	"encoding/json"
	"fmt"
	"net"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// 登录函数
func login(id int, password string) error {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		return err
	}
	defer conn.Close()

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
	// 将序列化好的消息给msg
	msg.Data = string(data)

	// 将msg序列化，然后发送给服务器
	data, err = json.Marshal(msg)
	if err != nil {
		return err
	}

	// 发送数据
	err = utils.WritePkg(conn, data)
	if err != nil {
		return err
	}

	// 获取数据
	msg, err = utils.ReadPkg(conn)
	if err != nil {
		return err
	}
	// 将msg的data部分反序列化
	var loginResult message.LoginResult
	err = json.Unmarshal([]byte(msg.Data), &loginResult)
	if err != nil {
		return err
	}
	if loginResult.Code == 200 {
		fmt.Println("登录成功")
	} else {
		fmt.Println(loginResult.Error)
	}

	return nil
}
