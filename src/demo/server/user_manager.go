package main

import (
	"fmt"
	"net"
)

var userManager *UserManager

type UserManager struct {
	OnlineUsers map[int]net.Conn
}

func init() {
	userManager = &UserManager{OnlineUsers: make(map[int]net.Conn)}
}

// AddOnlineUser 用户上线
func (this *UserManager) AddOnlineUser(userId int, conn net.Conn) {
	this.OnlineUsers[userId] = conn
}

// DeleteOnlineUser 用户下线
func (this *UserManager) DeleteOnlineUser(userId int) {
	delete(this.OnlineUsers, userId)
}

func (this *UserManager) GetOnlineUser(userId int) (net.Conn, error) {
	user, ok := this.OnlineUsers[userId]
	if ok == false {
		return nil, fmt.Errorf("user %d not exists or is not online\n", userId)
	}
	return user, nil
}

// GetAllOnlineUser 获取所有在线用户
func (this *UserManager) GetAllOnlineUser() map[int]net.Conn {
	return this.OnlineUsers
}
