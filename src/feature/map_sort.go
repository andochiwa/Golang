package main

import (
	"fmt"
	"sort"
)

func main() {
	mp := make(map[int]string)
	mp[2] = "h"
	mp[5] = "a"
	mp[0] = "c"
	mp[4] = "f"
	// go中对map排序需要借助切片来实现
	// 通过切片保存key，然后用sort包对key排序，再通过这些key回到原来的map中取值
	// smjb??????????????????????
	// 我要恶心吐了啊啊啊啊阿阿啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊
	var keys []int
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, mp[k])
	}
}
