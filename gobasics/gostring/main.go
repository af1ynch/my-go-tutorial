package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 字符串的基本特性
	name := "天下无双"

	fmt.Printf("字符串值：%s\n", name)
	fmt.Printf("字符串长度（字节）：%d\n", len(name))
	fmt.Printf("字符串大小（结构体）：%d 字节\n", unsafe.Sizeof(name))

	// 字符串创建

	// 方式三：从字节切片创建字符串
	bytes := []byte{72, 101, 108, 108, 111}
	fromBytes := string(bytes)
	fmt.Printf("字符串：%s\n", fromBytes)

	// 方式四：从 rune 切片创建字符串
	runes := []rune{'程', '序', '员', '许', '三', '多'}
	fromRunes := string(runes)
	fmt.Printf("字符串：%s\n", fromRunes)

}
