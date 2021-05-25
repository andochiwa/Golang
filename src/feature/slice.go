package main

import "fmt"

func changeSlice(slice []int) {
	slice[0] = 1
}

func main() {
	// 切片是数组的引用，因为传递时为引用传递
	arr := [...]int{3, 4, 5}
	// 声明一个切片，让切片引用数组的[1, 3)元素
	slice := arr[1:3]
	fmt.Println("1 arr =", arr)
	fmt.Println("1 slice =", slice)
	fmt.Println("1 slice元素个数 =", len(slice))
	fmt.Println("1 slice的容量 =", cap(slice))

	// 通过make函数声明切片 make(slice, size, [cap])
	slice = make([]int, 5, 10)
	fmt.Println("2 slice =", slice)
	fmt.Println("2 slice元素个数 =", len(slice))
	fmt.Println("2 slice的容量 =", cap(slice))
	changeSlice(slice)
	fmt.Println("2 change slice =", slice)

	// 通过直接指定具体数组来声明
	slice = []int{0, 0, 0}
	fmt.Println("3 slice =", slice)
	fmt.Println("3 slice元素个数 =", len(slice))
	fmt.Println("3 slice的容量 =", cap(slice))
	changeSlice(slice)
	fmt.Println("3 change slice =", slice)
}
