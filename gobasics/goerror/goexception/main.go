package main

import (
    "fmt"
    "runtime/debug"
)

// 面试鸭的题目处理系统
type QuestionProcessor struct {
    questions []string
    stats     map[string]int
}

func NewQuestionProcessor() *QuestionProcessor {
    return &QuestionProcessor{
        questions: make([]string, 0),
        stats:     make(map[string]int),
    }
}

// SafeProcessQuestion 安全的题目处理函数，使用 recover 捕获 panic
func (qp *QuestionProcessor) SafeProcessQuestion(questionID int, question string) (result string, err error) {
    // 使用 defer 和 recover 捕获可能的 panic
    defer func() {
        if r := recover(); r != nil {
            // 恢复 panic，转换为错误
            err = fmt.Errorf("题目处理发生严重错误: %v", r)

            // 记录 panic 的详细信息
            fmt.Printf("捕获到 panic: %v\n", r)
            fmt.Printf("堆栈信息:\n%s\n", debug.Stack())

            // 更新统计信息
            qp.stats["panic_count"]++
        }
    }()

    // 调用可能产生 panic 的函数
    return qp.processQuestionUnsafe(questionID, question), nil
}

// 可能产生 panic 的不安全函数
func (qp *QuestionProcessor) processQuestionUnsafe(questionID int, question string) string {
    // 更新统计
    qp.stats["total_processed"]++

    // 模拟各种可能导致 panic 的情况
    switch questionID {
    case 1:
        // 数组越界
        arr := []string{"a", "b", "c"}
        fmt.Println("访问数组元素:", arr[10]) // 这会导致 panic

    case 2:
        // 空指针解引用
        var ptr *string
        fmt.Println("解引用空指针:", *ptr) // 这会导致 panic

    case 3:
        // 除零错误
        zero := 0
        result := 100 / zero
        fmt.Println("除零结果:", result) // 这会导致 panic

    case 4:
        // 类型断言失败
        var x interface{} = "hello"
        num := x.(int) // 这会导致 panic
        fmt.Println("类型断言结果:", num)

    case 5:
        // 主动触发 panic
        panic("这是一个主动触发的 panic")

    default:
        // 正常处理
        qp.stats["success_count"]++
        return fmt.Sprintf("题目 %d 处理成功: %s", questionID, question)
    }

    return ""
}

// BatchProcess 批量处理题目，展示 panic 的传播和恢复
func (qp *QuestionProcessor) BatchProcess(questions map[int]string) {
    fmt.Println("=== 批量处理题目 ===")

    for questionID, question := range questions {
        fmt.Printf("\n处理题目 %d: %s\n", questionID, question)

        result, err := qp.SafeProcessQuestion(questionID, question)
        if err != nil {
            fmt.Printf("❌ 处理失败: %v\n", err)

            // 记录错误并继续处理其他题目
            qp.stats["error_count"]++
        } else {
            fmt.Printf("✅ 处理结果: %s\n", result)
        }
    }
}

// 演示 panic 的传播
func demonstratePanicPropagation() {
    fmt.Println("\n=== panic 传播演示 ===")

    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("在 main 函数中捕获到 panic: %v\n", r)
        }
    }()

    func() {
        defer func() {
            fmt.Println("内层函数的 defer 执行")
        }()

        fmt.Println("即将触发 panic...")
        panic("演示 panic 传播")

        fmt.Println("这行代码不会执行")
    }()

    fmt.Println("这行代码也不会执行")
}

func main() {
    fmt.Println("=== 面试鸭题目处理系统 ===")

    processor := NewQuestionProcessor()

    // 准备测试题目
    questions := map[int]string{
        0: "正常题目：什么是 Go 语言？",
        1: "数组越界测试题目",
        2: "空指针测试题目",
        3: "除零测试题目",
        4: "类型断言测试题目",
        5: "主动 panic 测试题目",
        6: "另一个正常题目：Go 语言的特点是什么？",
    }

    // 批量处理
    processor.BatchProcess(questions)

    // 显示统计信息
    fmt.Printf("\n=== 处理统计 ===\n")
    fmt.Printf("总处理数: %d\n", processor.stats["total_processed"])
    fmt.Printf("成功处理: %d\n", processor.stats["success_count"])
    fmt.Printf("发生 panic: %d\n", processor.stats["panic_count"])
    fmt.Printf("处理错误: %d\n", processor.stats["error_count"])

    // 演示 panic 传播
    demonstratePanicPropagation()

    fmt.Println("程序正常结束")
}
