package main

import (
	"fmt"
	"runtime"
	"time"
)

func cpuIntensiveTask(id int) {
	fmt.Printf("任务 %d 开始执行\n", id)

	count := 0
	for i := 1; i < 1000000000; i++ {
		count++

		// 每执行一定次数主动让出 CPU
		if i%100000000 == 0 {
			runtime.Gosched() // 主动让出执行权
			fmt.Printf("任务 %d 让出 CPU，已执行 %d 次\n", id, i)
		}
	}

	fmt.Printf("任务 %d 完成，总计算次数: %d\n", id, count)
}

func main() {
	fmt.Println("=== 协作式调度演示 ===")

	// 限制只使用一个逻辑处理器
	runtime.GOMAXPROCS(1)

	start := time.Now()

	// 启动两个CPU密集型任务
	go cpuIntensiveTask(1)
	go cpuIntensiveTask(2)

	// 等待任务完成
	time.Sleep(3 * time.Second)

	fmt.Printf("总耗时: %v\n", time.Since(start))
}
