package main

import (
	"fmt"
	"runtime"
)

// 剪切助手的数据处理系统
type DataProcessor struct {
	processedCount int
	errorCount     int
	panicCount     int
}

func NewDataProcessor() *DataProcessor {
	return &DataProcessor{}
}

// 安全的数据处理包装器
func (dp *DataProcessor) SafeProcessData(data interface{}, processor func(interface{}) interface{}) (result interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			dp.panicCount++

			// 获取 panic 的调用栈信息
			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			stackTrace := string(buf[:n])

			err = fmt.Errorf("数据处理发生 panic: %v\n调用栈:\n%s", r, stackTrace)

			fmt.Printf("⚠️  捕获到 panic: %v\n", r)
		}
	}()

	result = processor(data)
	dp.processedCount++
	return result, nil
}

// 各种数据处理函数（可能产生 panic）
func processString(data interface{}) interface{} {
	// 强制类型转换，可能 panic
	str := data.(string)
	return fmt.Sprintf("处理后的字符串: %s", str)
}

func processSlice(data interface{}) interface{} {
	// 切片操作，可能越界
	slice := data.([]int)
	return slice[0] + slice[len(slice)-1] // 如果切片为空会 panic
}

func processMap(data interface{}) interface{} {
	// Map 操作
	m := data.(map[string]int)
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func processWithPanic(data interface{}) interface{} {
	// 主动触发 panic
	panic("这是一个测试 panic")
}

func main() {
	fmt.Println("=== 剪切助手数据处理系统 ===")

	processor := NewDataProcessor()

	// 测试数据和处理函数
	testCases := []struct {
		name      string
		data      interface{}
		processor func(interface{}) interface{}
	}{
		{"字符串处理", "Hello, 剪切助手!", processString},
		{"错误类型转换", 123, processString}, // 会 panic
		{"切片处理", []int{1, 2, 3, 4, 5}, processSlice},
		{"空切片处理", []int{}, processSlice}, // 会 panic
		{"Map处理", map[string]int{"a": 1, "b": 2, "c": 3}, processMap},
		{"错误Map处理", "not a map", processMap},  // 会 panic
		{"主动Panic", "任意数据", processWithPanic}, // 会 panic
	}

	for i, tc := range testCases {
		fmt.Printf("\n--- 测试 %d: %s ---\n", i+1, tc.name)

		result, err := processor.SafeProcessData(tc.data, tc.processor)
		if err != nil {
			fmt.Printf("❌ 处理失败: %v\n", err)
			processor.errorCount++
		} else {
			fmt.Printf("✅ 处理成功: %v\n", result)
		}
	}

	// 统计信息
	fmt.Printf("\n=== 处理统计 ===\n")
	fmt.Printf("总测试数: %d\n", len(testCases))
	fmt.Printf("成功处理: %d\n", processor.processedCount)
	fmt.Printf("处理错误: %d\n", processor.errorCount)
	fmt.Printf("发生panic: %d\n", processor.panicCount)
}
