package main

import "fmt"

func main() {
	var arr [3]int = [3]int{2, 3, 8}
	arr[1] = 1
	fmt.Println(arr)

	var arr2 = [...]int{3, 4, 8}
	fmt.Println(arr2)

	var names = [...]string{1: "tom", 2: "jack", 0: "marry"}
	for i, name := range names {
		fmt.Println(i, name)
	}
}
