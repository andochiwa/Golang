package main

import "fmt"

func main() {
	// 声明map，只声明map不会分配内存，分配内存需要make
	var mp map[int]string
	mp = make(map[int]string)

	// CRUD
	// 插入
	mp[1] = "h"
	mp[2] = "a"
	mp[0] = "b"
	fmt.Println(mp)
	// 删除，使用内置函数delete
	delete(mp, 1)
	fmt.Println(mp)
	// 查找 用第二个返回值来判断是否存在
	val, find := mp[5]
	if find {
		fmt.Println(val)
	}

	// 遍历map
	for k, v := range mp {
		fmt.Println("k =", k, "v =", v)
	}
}
