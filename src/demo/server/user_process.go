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

	user, err := MyUserDao.Login(loginMessage.UserId, loginMessage.UserPwd)
	if err != nil {
		if err == ErrorUserNotExists {
			loginResult.Code = 444
		} else if err == ErrorUserPwd {
			loginResult.Code = 403
		}
		loginResult.Error = err.Error()
	} else {
		loginResult.Code = 200
		fmt.Println(user, "登录成功")
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

func ServerProcessRegister(conn net.Conn, mes *message.Message) error {
	// 取出message.data
	var user User
	var registerResult message.RegisterResult
	err := json.Unmarshal([]byte(mes.Data), &user)
	if err != nil {
		return err
	}
	err = MyUserDao.Register(&user)
	if err != nil {
		if err == ErrorUserExists {
			registerResult.Code = 444
		} else {
			registerResult.Code = 403
		}
		registerResult.Error = err.Error()
	} else {
		registerResult.Code = 200
		fmt.Println("注册成功")
	}
	// 把结果发送给客户端
	registerResultData, err := json.Marshal(registerResult)
	if err != nil {
		return err
	}
	resultMessage := message.Message{Type: message.RegisterResultType, Data: string(registerResultData)}
	resultMessageData, err := json.Marshal(resultMessage)
	if err != nil {
		return err
	}
	err = utils.WritePkg(conn, resultMessageData)
	if err != nil {
		return err
	}
	return nil

}
