package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello, %s! (%d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== 创建第一个 Goroutine ===")

	// 普通函数调用
	fmt.Println("普通调用:")
	sayHello("编程导航")

	fmt.Println("\n并发调用:")
	// 创建 Goroutine
	go sayHello("面试鸭")
	go sayHello("算法导航")

	// 主函数需要等待一段时间，否则程序会立即结束
	time.Sleep(500 * time.Millisecond)
	fmt.Println("主函数结束")
}
