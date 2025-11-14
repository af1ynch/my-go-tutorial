package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("=== Go 环境信息检查 ===")

	// 显示 Go 版本
	fmt.Printf("Go 版本：%s\n", runtime.Version())
	fmt.Printf("操作系统：%s\n", runtime.GOOS)
	fmt.Printf("架构：%s\n", runtime.GOARCH)

	// 显示环境变量
	fmt.Printf("GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("GOPATH: %s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOPROXY: %s\n", os.Getenv("GOPROXY"))

	// 测试基本功能
	fmt.Println("=== 基本功能测试 ===")
	testSlice()
	testMap()

	// 验证第三方包导入
	fmt.Println("=== 第三方包导入测试 ===")
	fmt.Printf("1和2是否相等：%t\n", cmp.Equal(1, 2))
}

func testSlice() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("切片测试：%v\n", nums)
}

func testMap() {
	m := map[string]int{"apple": 5, "banana": 3}
	fmt.Printf("映射测试：%v\n", m)
}
