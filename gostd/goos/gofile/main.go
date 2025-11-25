package main

import (
	"fmt"
	"os"
)

func main() {
	// 以只读方式打开文件
	file, err := os.Open("gostd/goos/gofile/test.txt")
	if err != nil {
		fmt.Printf("打开文件时发生错误：%v\n", err)
	}

	// 打开文件，如果文件已存在，则清空文件内容
	file1, err := os.Create("gostd/goos/gofile/test.txt")
	if err != nil {
		fmt.Printf("创建文件时发生错误：%v\n", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("file已被关闭")
		}

		if err := file1.Close(); err != nil {
			fmt.Printf("file1已被关闭")
		}
	}()

}
