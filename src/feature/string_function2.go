package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	// 将指定的子串替换成另一个子串 strings.Replace(string, oldstr, newstr, n) n指替换几个，-1为所有
	fmt.Println("hello world的前面两个l替换成j:", strings.Replace(str, "l", "j", 2))

	// 分割字符串 strings.Split(string, str)
	fmt.Println("分割字符串hello,world:", strings.Split("hello,world", ","))

	// 大小写转换 strings.ToLower(string), strings.ToUpper(string)
	fmt.Println("Go大写:", strings.ToUpper("Go"), "小写:", strings.ToLower("Go"))

	// 去除两边空格 			strings.TrimSpace(string)
	// 去除两边指定字符		strings.Trim(String, substr)
	// 去除左右边指定字符		strings.LeftTrim(string, substr)  strings.RightTrim(String, substr)

	// 判断是否以指定字符串开头	strings.HasPrefix(string, substr)
	// 判断是否以指定字符串结尾 strings.HasSuffix(string, substr)
}
