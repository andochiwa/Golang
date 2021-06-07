package main

var OnlineUsers = make(map[int]int)

func InsertUser(userId int) {
	OnlineUsers[userId] = 1
}

func deleteUser(userId int) {
	delete(OnlineUsers, userId)
}
