package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer conn.Close()

	fmt.Println("redis conn success")

	// 插入一条数据
	_, err = conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set error,", err)
		return
	}
}
