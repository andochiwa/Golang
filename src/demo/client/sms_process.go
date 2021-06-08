package main

import (
	"encoding/json"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

func SendMessage(content string) error {
	var msg message.Message
	msg.Type = message.SmsMessageType

	var smsMessage message.SmsMessage
	smsMessage.User = currentUser.User
	smsMessage.Content = content
	smsData, err := json.Marshal(smsMessage)
	if err != nil {
		return err
	}

	msg.Data = string(smsData)
	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = utils.WritePkg(currentUser.Conn, msgData)
	if err != nil {
		return err
	}
	return nil

}
