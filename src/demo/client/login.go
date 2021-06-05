package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"redis.demo/common/message"
	"time"
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

	// 先发送消息的长度
	// 转成byte
	pkgLen := uint32(len(data))
	// 因为uint32是4字节，所以只需要开4字节的byte
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:], pkgLen)
	// 发送长度
	n, err := conn.Write(bytes[:])
	if err != nil {
		return err
	} else if n != 4 {
		return errors.New("conn.Write send byte error")
	}

	// 发送消息
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	time.Sleep(time.Second * 10)
	return nil
}
