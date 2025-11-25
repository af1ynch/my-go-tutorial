package main

import (
	"fmt"
	"strings"
)

func main() {
	// 从字符串中读取数据
	reader := strings.NewReader("欢迎来到编程世界！")

	buffer := make([]byte, 12)

	n, err := reader.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("读取了 %d 字节，内容：%s\n", n, buffer)

}
