package main

import "fmt"

func main() {
	// 声明一个map切片
	mp := make([]map[int]string, 2)
	if mp[0] == nil {
		mp[0] = make(map[int]string)
		mp[0][0] = "hello"
		mp[0][1] = "world"
	}
	if mp[1] == nil {
		mp[1] = make(map[int]string)
		mp[1][0] = "HELLO"
		mp[1][1] = "WORLD"
	}
	// 动态增加一个map
	mp = append(mp, map[int]string{
		1: "HeLLo",
		2: "WoRlD",
	})
	fmt.Println(mp)
}
