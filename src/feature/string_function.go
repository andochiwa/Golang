package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串常用系统函数
func main() {
	str := "hello world"
	// len(string) 统计字符串长度
	fmt.Println("str len =", len(str))

	// []rune(string) 转成切片，可遍历出中文字符
	str = "北京"
	temp := []rune(str)
	for i := 0; i < len(temp); i++ {
		fmt.Printf("%c", temp[i])
	}
	fmt.Println()

	// 字符串转整数 strconv.Atoi(string)
	n, _ := strconv.Atoi("64")
	fmt.Println("字符串64转整数 =", n)
	// 整数转字符串 strconv.Itoa(int)
	str = strconv.Itoa(123)
	fmt.Println("整数123转字符串 =", str)
	// 字符串转byte []byte(string), byte转字符串 string([]byte)
	bytes := []byte(str)
	fmt.Println("字符串转byte =", bytes)
	// 10进制转2, 8, 16进制, strconv.FormatInt(int64, int)
	str = strconv.FormatInt(64, 2)
	fmt.Println("64转二进制 =", str)

	// 匹配子串 strings.Contains(string, string) Rabin-Karp算法
	fmt.Println("qwerty是否包含子串wer:", strings.Contains("qwerty", "wer"), "qwerty是否包含子串wet:", strings.Contains("qwerty", "wet"))
	// 统计子串个数 string.Count(String, String)
	fmt.Println("kksljfikkkkslkksfkksfwakk中kk的个数:", strings.Count("kksljfikkkkslkksfkksfwakk", "kk"))

	// 不区分大小写比较字符串 strings.EqualsFold(String, string)
	fmt.Println("abc = ABC ?", strings.EqualFold("abc", "ABC"))

	// 返回子串在主串中出现的位置 strings.Index(string, string)
	fmt.Println("efg在abc_efggd出现的位置:", strings.Index("abc_efggd", "efg"))
}
