package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 切片的定义与初始化

	resumeIDs := make([]int, 3, 5)
	resumeIDs[0] = 1001
	resumeIDs[1] = 1002
	resumeIDs[2] = 1003

	resumeIDs = append(resumeIDs, 1004)
	fmt.Printf("简历ID列表：%v\n", resumeIDs)

	arr := [5]string{"技术", "设计", "销售", "管理", "教育"}
	sli := []string{"技术", "设计", "销售", "管理", "教育"}

	fmt.Println("\n === 占用内存对比 === ")
	fmt.Printf("数组占用内存：%d 字节\n", reflect.TypeOf(arr).Size())
	fmt.Printf("切片占用内存：%d 字节\n", reflect.TypeOf(sli).Size())
}
