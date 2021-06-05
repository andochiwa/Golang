package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"redis.demo/common/message"
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
	err = writePkg(conn, data)
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	pkgLen := uint32(len(data))

	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:], pkgLen)
	// 发送长度
	n, err := conn.Write(bytes[:])
	if err != nil {
		return
	} else if n != 4 {
		err = errors.New("conn.Write send byte error")
		return
	}

	// 发送消息
	_, err = conn.Write(data)
	if err != nil {
		return
	}
	return
}
