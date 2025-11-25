package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go 调度器信息 ===")

	// 获取系统信息
	fmt.Printf("CPU 核心数: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("当前 Goroutine 数: %d\n", runtime.NumGoroutine())

	// 模拟面试鸭并发处理面试题
	fmt.Println("\n=== 面试鸭题库并发处理 ===")

	var wg sync.WaitGroup
	questionCount := 20

	// 处理题目的函数
	processQuestion := func(questionID int) {
		defer wg.Done()

		// 模拟处理时间
		processingTime := time.Duration(questionID%3+1) * 100 * time.Millisecond

		fmt.Printf("开始处理题目 %d (预计耗时: %v)\n", questionID, processingTime)
		time.Sleep(processingTime)
		fmt.Printf("完成处理题目 %d\n", questionID)
	}

	start := time.Now()

	// 启动多个 Goroutine 处理题目
	for i := 1; i <= questionCount; i++ {
		wg.Add(1)
		go processQuestion(i)
	}
	fmt.Printf("运行时 Goroutine 数量： %d\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Printf("处理完成，总耗时: %v\n", time.Since(start))
	fmt.Printf("最终 Goroutine 数: %d\n", runtime.NumGoroutine())

	// go spinner(100 * time.Millisecond)
	// fmt.Printf("运行时 Goroutine 数量： %d\n", runtime.NumGoroutine())
	//
	// const n = 45
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	fmt.Println("=== 主协程与子协程关系演示 === ")

	fmt.Println("启动子协程")

	// 启动三个子协程
	for i := 0; i < 3; i++ {
		go childGoroutine(i)
	}

	fmt.Println("主协程继续执行...")

	// 主协程只等1秒就结束
	time.Sleep(time.Second)

	fmt.Println("主协程结束，程序退出")

	// 演示正确等待子Goroutine
	fmt.Println("=== 演示正确等待子Goroutine执行 === ")

	var wg1 sync.WaitGroup

	userIDs := []int{1001, 1002, 1003, 1004, 1005}
	fmt.Printf("需要处理%d个用户的数据\n", len(userIDs))

	start1 := time.Now()

	// 为每个用户启动一个协程
	for _, userID := range userIDs {
		wg1.Add(1) // 增加等待的协程数量
		go processUserData(userID, &wg1)
	}

	fmt.Println("所有处理协程已启动，等待执行完成...")
	// 等待所有协程完成
	wg1.Wait()

	fmt.Printf("所有用户数据处理完成，总耗时：%v\n", time.Since(start1))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func childGoroutine(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("子协程 %d；第%d次执行\n", id, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func processUserData(userID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("开始处理用户%d的数据\n", userID)

	// 模拟数据处理
	time.Sleep(time.Duration(userID) * 200 * time.Millisecond)

	fmt.Printf("用户%d的数据处理完成")
}
