/*
	Golang基础语法

1. Go程序结构
1.1. 程序执行流程
2. 包的概念与使用
2.1. 包的基本概念
2.2. 标准库包的使用
2.3. 包的导入方式
2.4. 自定义包
3. 变量的声明与初始化
3.1. 变量声明的四种方式

	标准
	类型推断
	短变量声明（函数内）
	批量声明

3.2. 变量的零值
3.3. 多变量声明
3.4. 变量的作用域
4. 常量的定义与使用
4.1. 常量的基本定义
4.2. 常量的特点
4.3. iota
4.4. 常量的实际应用
5. 基本数据类型
5.1. 整数类型
5.2. 浮点数类型
5.3. 字符类型
5.4. 字符串类型
5.5. 布尔类型
6. 类型转换与类型推断
6.1. 基本类型转换
6.2. 类型推断
6.3. 类型断言
7. Go 语言命名规范
7.1. 基本命名规则
7.2. 常量命名
7.3. 特殊命名约定
7.4. 命名最佳实践
*/
package main

import (
	"fmt"    // 格式化输入输出
	"math"   // 数学函数
	m "math" // 给包起别名
	"math/rand"
	"strconv" // 字符串转换
	"strings" // 字符串处理
	"time"    // 时间处理

	. "strings" // 点导入，可以直接使用包内的函数
	// _ "image/png"  // 匿名导入，只执行包的 init 函数

	"github.com/af1ynch/my-go-tutorial/caculator"
)

const (
	MONDAY = iota + 1
	TUESDAY
	WENSDAY
	THURSDAY
	FRIDAY
	STATURDAY
	SUNDAY
)

func main() {

	// 包的概念与使用
	// 使用 fmt 包
	fmt.Println("Hello, World")

	// 使用 math 包
	fmt.Printf("圆周率：%.2f\n", math.Pi)
	fmt.Printf("2的10次方：%.0f\n", math.Pow(2, 10))

	// 使用包别名
	fmt.Printf("正弦值：%.2f\n", m.Sin(m.Pi/2))

	// 使用子包
	fmt.Printf("随机数：%d\n", rand.Intn(100))

	// 点导入允许直接使用函数名
	fmt.Printf("转换为大写：%s\n", ToUpper("Hello"))

	// 使用 strings 包
	text := "Hello, Go Programming!"
	fmt.Printf("转换为大写：%s\n", strings.ToUpper(text))
	fmt.Printf("包含 Go：%t\n", strings.Contains(text, "Go"))

	// 使用 time 包
	now := time.Now()
	fmt.Printf("当前时间：%v\n", now)
	fmt.Printf("当前时间：%s\n", now.Format("2006-01-02 15:04:05"))

	//  使用 strconv 包
	num := "123"
	value, err := strconv.Atoi(num)
	if err == nil {
		fmt.Printf("字符串'%s'转换为整数：%d\n", num, value)
	}

	// 使用自定义包
	fmt.Printf("半径为 5 的圆面积是：%.2f\n", caculator.CircleArea(5))

	// 变量

	// 变量声明
	var var1 string // 标准
	var i = 10      // 类型推断
	g := "Hello."

	fmt.Printf("字符串的零值是：%v\n", var1)
	fmt.Printf("数：%d\n", i)
	fmt.Printf("你好：%s\n", g)

	// 基本数据类型
	gr := "Hello"
	fmt.Printf("g: %+v\n", gr)

	// 字符类型
	b1 := 'A'
	fmt.Printf("byte: %c, 数值：%d\n", b1, b1)
	r1 := '中'
	fmt.Printf("byte: %c, 数值: %d\n", r1, r1)
	// 字符串

	// 字符串拼接

	// 方式一：格式化
	name := "张三"
	age := 25
	msg := fmt.Sprintf("你好，我是%s, 今年%d岁", name, age)
	fmt.Printf("方式一：格式化 - %s\n", msg)

	// 方式二：+
	msg2 := "你好，我是" + name + "，今年" + strconv.Itoa(age) + "岁"
	fmt.Printf("方式二：+操作符 - %s\n", msg2)

	// 字符串替换
	s1 := "Hello, 世界，世界，世界...界..."
	replaced := strings.Replace(s1, "世界", "Go", 1)
	fmt.Printf("只替换一个字符串：%s\n", replaced)
	replaced2 := strings.ReplaceAll(s1, "世界", "Go")
	fmt.Printf("替换所有字符串：%s\n", replaced2)

	// 字符串遍历

	// 按字节遍历
	fmt.Print("按字节遍历：")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i])
	}
	fmt.Println()
	// 按字符遍历
	fmt.Print("按字符遍历：")
	for _, r := range s1 {
		fmt.Printf("%c", r)
	}
	fmt.Println()
	// 类型转换
	// 浮点数转换为整数
	pi := 3.14159
	tentimesPi := 3141.59
	iPi := int(pi)
	fmt.Printf("pi: %.5f, int(pi): %d\n", pi, iPi)

	// 浮点转换为字符串

	// 精细控制：strconv.FormatFloat
	fmt.Printf("pi: %.5f, stringf(pi): %s\n", pi, strconv.FormatFloat(pi, 'f', 5, 64))
	fmt.Printf("pi: %.5f, stringe(pi): %s\n", pi, strconv.FormatFloat(tentimesPi, 'e', 5, 64))
	fmt.Printf("pi: %.5f, stringg(pi): %s\n", pi, strconv.FormatFloat(pi, 'g', 5, 64))

	// 简单的格式化
	fmt.Println(fmt.Sprintf("%f", pi))
	fmt.Println(fmt.Sprintf("%.5f", pi))
	fmt.Println(fmt.Sprintf("%e", pi))

	fmt.Printf("星期一：%d\n", MONDAY)
	fmt.Printf("星期二：%d\n", TUESDAY)
	fmt.Printf("星期三：%d\n", WENSDAY)
	fmt.Printf("星期四：%d\n", THURSDAY)
	fmt.Printf("星期五：%d\n", FRIDAY)
	fmt.Printf("星期六：%d\n", STATURDAY)
	fmt.Printf("星期日：%d\n", SUNDAY)

	var sm = map[string]float64{
		"Go":     95.0,
		"Python": 99.0,
		"MySQL":  100.0,
	}
	fmt.Printf("平均分是：%.2f\n", average(sm))

}

func average(scoreMap map[string]float64) float64 {
	var sum float64
	for _, val := range scoreMap {
		sum += val
	}
	return sum / float64(len(scoreMap))
}
