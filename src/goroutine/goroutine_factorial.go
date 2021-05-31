package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

var (
	mp   = make(map[int]int64, 0)
	lock = sync.Mutex{}
)

// 计算1-200的各个数的总和，放到map中，要求使用协程并行计算，处理协程安全问题
// 有三种方法处理map协程安全问题，第一种是使用sync.Map
// 第二种是使用互斥锁sync.Mutex，第三种是使用管道channel
func factorial(n int) {
	var res int64 = 0
	for i := 1; i <= n; i++ {
		res += int64(i)
	}
	lock.Lock()
	mp[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go factorial(i)
	}
	runtime.Gosched()

	var keys []int
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("阶乘 %v = %v\n", k, mp[k])
	}
}
