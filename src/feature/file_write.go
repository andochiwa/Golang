package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "src/file/test2.txt"
	// 只读 如果不存在则创建
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// 写入到缓存
		_, _ = writer.WriteString(str)
	}
	// 写入到文件
	_ = writer.Flush()
}
