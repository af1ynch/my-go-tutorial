/* 练习：
1. 个人信息管理系统
  创建一个简单的个人信息管理系统，实现以下功能：
  1.1. 定义用户信息结构（姓名、年龄、邮箱、电话）
  1.2. 实现信息的输入、显式和验证
  1.3. 使用适当的数据类型和命名规范
2. 数据类型转换器
  创建一个数据类型转换工具，支持常见类型之间的转换
  2.1. 支持字符串到各种数字类型的转换
  2.2. 支持数字类型到字符串的转换
  2.3. 提供错误处理和类型检查
3. 命名规范检查器
  创建一个工具来检查代码中的命名是否符合 Go 语言规范
  3.1. 检查变量、函数、常量的命名规范
  3.2. 提供修正建议
  3.3. 支持批量检查
*/

package main

import (
    "fmt"
    "regexp"
    "strings"
)

type User struct {
    name     string
    age      int
    email    string
    phone    string
    isActive bool
}

// 验证邮箱格式
func isValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z{2,}$]`)
    return emailRegex.MatchString(email)
}

// 验证手机号格式
func isValidPhone(phone string) bool {
    phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
    return phoneRegex.MatchString(phone)
}

// 验证年龄
func isValidAge(age int) bool {
    return age > 0 && age < 200
}

// CreateUser 创建用户信息
func CreateUser(name, email, phone string, age int) (*User, error) {
    if strings.TrimSpace(name) == "" {
        return nil, fmt.Errorf("姓名不能为空")
    }

    if !isValidEmail(email) {
        return nil, fmt.Errorf("邮箱格式错误")
    }

    if !isValidPhone(phone) {
        return nil, fmt.Errorf("手机号码格式错误")
    }

    if !isValidAge(age) {
        return nil, fmt.Errorf("年龄必须在1~200之间")
    }

    return &User{
        name:     name,
        age:      age,
        email:    email,
        phone:    phone,
        isActive: true,
    }, nil
}

// Display 显式用户信息
func (u *User) Display() {
    fmt.Println("=== 个人信息 ===")
    fmt.Printf("姓名：%s\n", u.name)
    fmt.Printf("年龄：%d\n", u.age)
    fmt.Printf("邮箱：%s\n", u.email)
    fmt.Printf("手机号码：%s\n", u.phone)
    fmt.Printf("状态：%s\n", map[bool]string{true: "活跃", false: "不活跃"}[u.isActive])
}

// UpdateAge 更新年龄
func (u *User) UpdateAge(newAge int) error {
    if !isValidAge(newAge) {
        return fmt.Errorf("年龄必须在1-200之间")
    }
    u.age = newAge
    return nil
}

func main() {
    // 创建用户信息
    user1, err := CreateUser("张三", "admin@123.com", "13912200033", 30)
    if err != nil {
        fmt.Printf("创建用户信息失败：%v\n", err)
        return
    }

    // 显式用户信息
    user1.Display()

    // 更新年龄
    err = user1.UpdateAge(31)
    if err != nil {
        fmt.Printf("更新年龄失败：%v\n", err)
    } else {
        fmt.Printf("年龄更新成功，新年龄：%d\n", user1.age)
    }

    // 测试无效数据
    fmt.Println("\n测试无效数据：")
    invalidCases := []struct {
        name, email, phone string
        age                int
    }{
        {"", "test@example.com", "13577799222", 22},
        {"张三", "invalid-email@1", "13477700222", 30},
        {"李四", "test@123.com", "01013222000", 35},
        {"王二麻", "admin@123.com", "13516300352", 201},
    }

    for i, testCase := range invalidCases {
        _, err := CreateUser(testCase.name, testCase.email, testCase.phone, testCase.age)
        if err != nil {
            fmt.Printf("测试案例：%d：%v\n", i+1, err)
        }
    }
}
