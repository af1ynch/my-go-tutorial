package main

import (
	"fmt"
	"runtime"
	"time"
)

// 性能测试函数
func performanceTest(bufferSize int, dataCount int) time.Duration {
	dataChan := make(chan int, bufferSize)
	start := time.Now()

	// 生产者
	go func() {
		for i := 0; i < dataCount; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	// 消费者
	count := 0
	for range dataChan {
		count++
	}

	return time.Since(start)
}

func main() {
	fmt.Println("=== Channel 缓冲区大小性能对比 ===")

	dataCount := 10000
	bufferSizes := []int{0, 1, 10, 100, 1000}

	fmt.Printf("处理 %d 个数据项的性能对比:\n", dataCount)

	for _, bufferSize := range bufferSizes {
		duration := performanceTest(bufferSize, dataCount)
		bufferType := "无缓冲"
		if bufferSize > 0 {
			bufferType = fmt.Sprintf("%d 缓冲", bufferSize)
		}

		fmt.Printf("%s Channel: %v\n", bufferType, duration)
	}

	// 内存使用情况演示
	fmt.Println("\n=== 内存使用情况演示 ===")

	runtime.GC() // 强制垃圾回收
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 创建大缓冲区 Channel
	largeChan := make(chan [1024]byte, 1000) // 1000个1KB的数据

	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("创建大缓冲区前内存: %d KB\n", m1.Alloc/1024)
	fmt.Printf("创建大缓冲区后内存: %d KB\n", m2.Alloc/1024)
	fmt.Printf("内存增加: %d KB\n", (m2.Alloc-m1.Alloc)/1024)

	// 释放 Channel
	a := [1024]byte{'a'}
	largeChan <- a
	largeChan = nil
	runtime.GC()
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)

	fmt.Printf("释放缓冲区后内存: %d KB\n", m3.Alloc/1024)
}
