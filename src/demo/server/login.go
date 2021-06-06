package main

import (
	"encoding/json"
	"fmt"
	"net"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// ServerProcessLogin 处理登录流程
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 取出mes.data 并反序列化
	var loginMessage message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data), &loginMessage)
	if err != nil {
		return
	}

	var resultMessage message.Message
	resultMessage.Type = message.LoginResultType
	var loginResult message.LoginResult

	// 如果id为100，密码为abc就认为合法
	fmt.Println("message =", loginMessage)
	if loginMessage.UserId == 100 && loginMessage.UserPwd == "abc" {
		loginResult.Code = 200
	} else {
		loginResult.Code = 444
		loginResult.Error = "User not exist"
	}

	// 把loginResult序列化
	data, err := json.Marshal(loginResult)
	if err != nil {
		return
	}
	resultMessage.Data = string(data)
	// 把result序列化，发给客户端
	data, err = json.Marshal(resultMessage)
	if err != nil {
		return
	}
	// 发送消息
	err = utils.WritePkg(conn, data)
	return
}
