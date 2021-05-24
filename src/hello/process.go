package main

import "fmt"

func main() {
	// *****************if********************
	var age int
	fmt.Println("输入年龄")
	fmt.Scanln(&age)
	if age > 20 {
		fmt.Println("if 年龄大于20")
	} else {
		fmt.Println("if 年龄小于20")
	}
	// golang支持在if中定义变量
	if name := "test"; name == "test" {
		fmt.Println("if name =", name)
	}

	// *****************switch********************
	// golang的switch不需要break
	switch {
	case age > 20:
		fmt.Println("switch 年龄大于20")
	default:
		fmt.Println("switch 年龄小于20")
	}

	// *****************for********************
	for i := 1; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// 遍历string
	str := "hello world"
	for idx, val := range str {
		fmt.Printf("%d:%c  ", idx, val)
	}
}
