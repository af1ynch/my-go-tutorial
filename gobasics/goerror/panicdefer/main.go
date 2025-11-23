package main

import (
	"fmt"
	"time"
)

func main() {
	demonstrateExecutionOrder()
	demonstrateMultipleRecover()
	demonstratePanicInDefer()
}

// defer、panic、recvoer执行顺序演示
func demonstrateExecutionOrder() {

	defer func() {
		fmt.Println("5. 最外层defer执行")
		if r := recover(); r != nil {
			fmt.Printf("6. 捕获到 panic：%v\n", r)
		}
		fmt.Println("7. 最外层defer完成")
	}()

	defer func() {
		fmt.Println("4. 外层defer执行")
	}()

	func() {
		defer func() {
			fmt.Println("3. 内层defer执行")
		}()

		fmt.Println("1. 开始处理简历")
		fmt.Println("2. 即将发生 panic")
		panic(fmt.Sprintf("简历格式错误：%v", time.Now()))
		fmt.Println("这个语句也不会执行")
	}()

	fmt.Println("这个语句不会执行")
}

// 多重defer和recover的处理
func demonstrateMultipleRecover() {

	defer func() {
		fmt.Println("外层defer开始")
		if r := recover(); r != nil {
			fmt.Printf("外层捕获：%v\n", r)
		}
		fmt.Println("外层defer结束")
	}()

	defer func() {
		fmt.Println("内层defer开始")
		if r := recover(); r != nil {
			fmt.Printf("内层捕获：%v\n", r)
		}
		panic("内层defer发生panic")
		fmt.Println("内层defer结束")
	}()

	panic("测试多重 recover")
}

// panic在defer中的传播
func demonstratePanicInDefer() {

	defer func() {
		fmt.Println("第三个 defer 执行")
		// if r := recover(); r != nil {
		// 	fmt.Printf("最终捕获：%v\n", r)
		// }
	}()

	defer func() {
		fmt.Println("第二个 defer 执行")
		if r := recover(); r != nil {
			fmt.Printf("%v\n", r)
		}
		panic("defer中的panic") // 会覆盖原来的panic
	}()

	defer func() {
		fmt.Println("第一个defer执行")
	}()

	panic("原始 panic")
}
