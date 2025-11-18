package main

import (
    "errors"
    "fmt"
    "os"
    "strconv"
    "time"
)

// 模拟用户分数验证
func validateScore(scoreStr string) (int, error) {
    // 尝试将字符串转换为整数
    score, err := strconv.Atoi(scoreStr)
    if err != nil {
        return 0, fmt.Errorf("转换错误：%v", err)
    }

    if score < 0 || score > 100 {
        return 0, fmt.Errorf("分数超出范围：%d，有效范围是0~100", score)
    }
    return score, nil
}

// 编程导航的课程错误类型
type CourseError struct {
    CourseID int
    Message  string
    Code     string
}

func (ce *CourseError) Error() string {
    return fmt.Sprintf("课程错误 [%s]: 课程ID %d - %s", ce.Code, ce.CourseID, ce.Message)
}

// 用户权限错误类型
type PermissionError struct {
    UserID       int
    Operation    string
    RequiredRole string
}

func (pe *PermissionError) Error() string {
    return fmt.Sprintf("权限不足: 用户 %d 尝试执行 '%s'，需要 %s 权限",
        pe.UserID, pe.Operation, pe.RequiredRole)
}

// 模拟课程访问检查
func checkCourseAccess(userID, courseID int, userRole string) error {
    // 检查课程是否存在
    if courseID <= 0 || courseID > 1000 {
        return &CourseError{
            CourseID: courseID,
            Message:  "课程不存在",
            Code:     "COURSE_NOT_FOUND",
        }
    }

    // 检查付费课程权限
    if courseID > 100 && userRole != "premium" {
        return &PermissionError{
            UserID:       userID,
            Operation:    "访问付费课程",
            RequiredRole: "premium",
        }
    }

    return nil
}

// 算法导航的文件操作示例
type FileManager struct {
    basePath string
}

func NewFileManager(basePath string) *FileManager {
    return &FileManager{basePath: basePath}
}

func (fm *FileManager) ReadFile(filename string) (string, error) {
    // 方式1: 使用 errors.New 创建简单错误
    if filename == "" {
        return "", errors.New("文件名不能为空")
    }

    // 方式2: 使用 fmt.Errorf 创建格式化错误
    if len(filename) > 100 {
        return "", fmt.Errorf("文件名过长: %d 字符，最大允许100字符", len(filename))
    }

    fullPath := fm.basePath + "/" + filename

    // 模拟文件读取
    if filename == "non_existent.txt" {
        return "", fmt.Errorf("文件不存在: %s", fullPath)
    }

    if filename == "permission_denied.txt" {
        return "", &os.PathError{
            Op:   "open",
            Path: fullPath,
            Err:  errors.New("permission denied"),
        }
    }

    // 模拟成功读取
    return fmt.Sprintf("算法导航文件内容: %s (读取时间: %s)",
        filename, time.Now().Format("15:04:05")), nil
}

func (fm *FileManager) WriteFile(filename, content string) error {
    if filename == "" {
        return errors.New("文件名不能为空")
    }

    if content == "" {
        return errors.New("文件内容不能为空")
    }

    // 模拟磁盘空间不足
    if len(content) > 1000 {
        return fmt.Errorf("磁盘空间不足: 需要 %d 字节，可用空间不足", len(content))
    }

    fmt.Printf("文件写入成功: %s (%d 字节)\n", filename, len(content))
    return nil
}

// ProcessFiles 批量处理文件
func (fm *FileManager) ProcessFiles(operations []struct {
    op       string
    filename string
    content  string
}) []error {
    var errs []error

    for i, operation := range operations {
        fmt.Printf("操作 %d: %s %s\n", i+1, operation.op, operation.filename)

        var err error
        switch operation.op {
        case "read":
            _, err = fm.ReadFile(operation.filename)
        case "write":
            err = fm.WriteFile(operation.filename, operation.content)
        default:
            err = fmt.Errorf("不支持的操作: %s", operation.op)
        }

        if err != nil {
            fmt.Printf("  ❌ 失败: %v\n", err)
            errs = append(errs, err)
        } else {
            fmt.Printf("  ✅ 成功\n")
            errs = append(errs, nil)
        }
    }

    return errs
}

// 剪切助手的数据同步系统
type SyncError struct {
    Operation string
    DeviceID  string
    Timestamp string
    Cause     error
}

func (se *SyncError) Error() string {
    return fmt.Sprintf("同步失败 [%s] 设备 %s 在 %s: %v",
        se.Operation, se.DeviceID, se.Timestamp, se.Cause)
}

func (se *SyncError) Unwrap() error {
    return se.Cause
}

// 网络连接错误
var (
    ErrNetworkTimeout  = errors.New("网络连接超时")
    ErrNetworkRefused  = errors.New("连接被拒绝")
    ErrInvalidResponse = errors.New("服务器响应无效")
    ErrAuthFailed      = errors.New("身份验证失败")
)

// 模拟网络请求
func performNetworkRequest(deviceID string) error {
    // 模拟不同的网络错误
    switch deviceID {
    case "device_001":
        return ErrNetworkTimeout
    case "device_002":
        return ErrAuthFailed
    case "device_003":
        return ErrInvalidResponse
    default:
        return nil
    }
}

// 同步数据到服务器
func syncToServer(deviceID, operation string) error {
    // 尝试网络请求
    if err := performNetworkRequest(deviceID); err != nil {
        // 包装错误，添加上下文信息
        return &SyncError{
            Operation: operation,
            DeviceID:  deviceID,
            Timestamp: "2024-01-15 10:30:00",
            Cause:     fmt.Errorf("网络请求失败: %w", err),
        }
    }

    return nil
}

// 批量同步
func batchSync(devices []string) error {
    var syncErrors []error

    for _, deviceID := range devices {
        if err := syncToServer(deviceID, "数据同步"); err != nil {
            syncErrors = append(syncErrors, err)
        }
    }

    if len(syncErrors) > 0 {
        return fmt.Errorf("批量同步失败: %d 个设备同步出错", len(syncErrors))
    }

    return nil
}

// 错误检查和处理
func handleSyncError(err error) {
    if err == nil {
        return
    }

    fmt.Printf("处理错误: %v\n", err)

    // 检查是否是特定的错误类型
    var syncErr *SyncError
    if errors.As(err, &syncErr) {
        fmt.Printf("  -> 同步错误详情: 操作=%s, 设备=%s, 时间=%s\n",
            syncErr.Operation, syncErr.DeviceID, syncErr.Timestamp)

        // 检查根本原因
        if errors.Is(syncErr.Cause, ErrNetworkTimeout) {
            fmt.Printf("  -> 建议: 检查网络连接，稍后重试\n")
        } else if errors.Is(syncErr.Cause, ErrAuthFailed) {
            fmt.Printf("  -> 建议: 重新登录或检查账户权限\n")
        } else if errors.Is(syncErr.Cause, ErrInvalidResponse) {
            fmt.Printf("  -> 建议: 检查服务器状态或联系技术支持\n")
        }
    }

    fmt.Println()
}

func main() {
    userScores := []string{"100", "50", "a10", "-50", "-0", "101"}

    for _, uscore := range userScores {
        score, err := validateScore(uscore)
        if err != nil {
            fmt.Printf("❌ 输入 '%s': %v\n", uscore, err)
        } else {
            fmt.Printf("✅ 输入 '%s'：%d\n", uscore, score)
        }
    }

    fmt.Println("=== 编程导航课程访问检查 ===")

    testCases := []struct {
        userID   int
        courseID int
        userRole string
    }{
        {1001, 50, "free"},      // 正常访问免费课程
        {1002, 200, "free"},     // 免费用户访问付费课程
        {1003, 1500, "premium"}, // 访问不存在的课程
        {1004, 200, "premium"},  // 付费用户正常访问
    }

    for _, tc := range testCases {
        fmt.Printf("用户 %d (角色: %s) 访问课程 %d: ", tc.userID, tc.userRole, tc.courseID)

        if err := checkCourseAccess(tc.userID, tc.courseID, tc.userRole); err != nil {
            fmt.Printf("❌ %v\n", err)

            // 根据错误类型进行不同处理
            var courseError *CourseError
            var permissionError *PermissionError
            switch {
            case errors.As(err, &courseError):
                fmt.Printf("   -> 建议：检查课程ID或联系管理员\n")
            case errors.As(err, &permissionError):
                fmt.Printf("   -> 建议：升级到付费会员\n")
            }

        }
        fmt.Println()
    }

    fmt.Println("=== 算法导航文件管理系统 ===")

    manager := NewFileManager("/data/algorithms")

    // 测试各种操作
    operations := []struct {
        op       string
        filename string
        content  string
    }{
        {"read", "bubble_sort.go", ""},
        {"read", "", ""},                      // 空文件名
        {"read", "non_existent.txt", ""},      // 不存在的文件
        {"read", "permission_denied.txt", ""}, // 权限错误
        {"write", "quick_sort.go", "package main\n\nfunc quickSort() {\n  // 快速排序实现\n}"},
        {"write", "large_file.txt", string(make([]byte, 2000))}, // 文件过大
        {"write", "", "content"},   // 空文件名
        {"delete", "test.txt", ""}, // 不支持的操作
    }

    allErrors := manager.ProcessFiles(operations)

    // 统计结果
    successCount := 0
    for _, err := range allErrors {
        if err == nil {
            successCount++
        }
    }

    fmt.Printf("\n=== 处理结果统计 ===\n")
    fmt.Printf("总操作数: %d\n", len(operations))
    fmt.Printf("成功: %d\n", successCount)
    fmt.Printf("失败: %d\n", len(operations)-successCount)
    fmt.Printf("成功率: %.1f%%\n", float64(successCount)/float64(len(operations))*100)

    fmt.Println("=== 剪切助手数据同步系统 ===")

    devices := []string{"device_001", "device_002", "device_003", "device_004"}

    // 单独同步每个设备
    fmt.Println("单独同步测试:")
    for _, deviceID := range devices {
        fmt.Printf("同步设备 %s: ", deviceID)

        err := syncToServer(deviceID, "数据同步")
        if err != nil {
            fmt.Printf("❌ 失败\n")
            handleSyncError(err)
        } else {
            fmt.Printf("✅ 成功\n")
        }
    }

    // 批量同步
    fmt.Println("批量同步测试:")
    batchErr := batchSync(devices[:3]) // 同步前3个设备（包含错误）
    if batchErr != nil {
        fmt.Printf("❌ 批量同步失败: %v\n", batchErr)
    } else {
        fmt.Printf("✅ 批量同步成功\n")
    }
}
