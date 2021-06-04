package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的连接池
var pool *redis.Pool

// 初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     10,  // 最大空闲连接数
		MaxActive:   0,   // 和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		// 初始化连接
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 从连接池取出连接
	conn := pool.Get()
	defer conn.Close()

	// 放入
	_, err := conn.Do("Set", "pool", "pool")
	if err != nil {
		return
	}

	// 取出
	str, err := conn.Do("Get", "pool")
	if err != nil {
		return
	}
	fmt.Printf("redis get key pool = %s", str)

}
