/*
	运算符与表达式

1. 算术运算符
1.1. 基本算术运算
1.2. 算术运算的特殊情况
1.3. 实际应用示例
2. 关系（比较）运算符
2.1. 基本比较运算
2.2. 浮点数比较的注意事项
2.3.
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	a, b := 11, 3

	// 基本算术运算
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("加法：%d + %d = %d\n", a, b, a+b)
	fmt.Printf("减法：%d - %d = %d\n", a, b, a-b)
	fmt.Printf("乘法：%d * %d = %d\n", a, b, a*b)
	fmt.Printf("除法：%d / %d = %d\n", a, b, a/b)
	fmt.Printf("取余：%d %% %d = %d\n", a, b, a%b)

	// 浮点数运算
	c, d := 11.0, 3.0
	fmt.Printf("\n浮点数运算：\n")
	fmt.Printf("%.1f / %.1f = %.3f\n", c, d, c/d) // 自动四舍五入

	// 一元运算符
	fmt.Printf("\n一元运算符：\n")
	fmt.Printf("+a = %d (二进制：%064b)\n", +a, +a)
	fmt.Printf("-b = %d (二进制：%064b)\n", -b, -b)

	// 算术运算的特殊情况

	// 整数除法
	fmt.Println("=== 整数除法 === ")
	fmt.Printf("7 / 3 = %d\n", 7/3) // 结果是2，不是2.333...，小数部分被截断
	fmt.Printf("7 %% 3 = %d\n", 7%3)

	// 浮点数除法
	fmt.Println("\n=== 浮点数除法 === ")
	fmt.Printf("7.0 / 3.0 = %.3f\n", 7.0/3.0)

	// 整数除零错误（编译时会自动检测）
	fmt.Println("10 / 0")

	// 浮点数除零特殊值

	// 除数与被除数都是浮点型字面量
	// fmt.Println("\n=== 浮点数特殊值 === ")
	// fmt.Printf("1.0 / 0.0 = %f\n", 1.0/0.0)   // 编译器报错
	// fmt.Printf("-1.0 / 0.0 = %f\n", -1.0/0.0) // 编译器报错
	// fmt.Printf("0.0 / 0.0 = %f\n", 0.0/0.0)   // 编译器报错

	// 除数与被除数都是浮点型变量
	e, f := 1.0, 0.0
	fmt.Println("\n=== 浮点数特殊值 === ")
	fmt.Printf("1.0 / 0.0 = %f\n", e/f)   // +inf
	fmt.Printf("-1.0 / 0.0 = %f\n", -e/f) // -inf
	fmt.Printf("0.0 / 0.0 = %f\n", f/f)   // NaN

	// 被除数是浮点型变量，除数是浮点型字面量
	fmt.Println("\n=== 浮点数特殊值 === ")
	fmt.Printf("1.0 / 0.0 = %f\n", e/0.0)   // +inf
	fmt.Printf("-1.0 / 0.0 = %f\n", -e/0.0) // -inf
	fmt.Printf("0.0 / 0.0 = %f\n", f/0.0)   // NaN

	// 被除数是浮点型字面量，除数是浮点型变量
	fmt.Println("\n=== 浮点数特殊值 === ")
	fmt.Printf("1.0 / 0.0 = %f\n", 1.0/f)   // +inf
	fmt.Printf("-1.0 / 0.0 = %f\n", -1.0/f) // -inf
	fmt.Printf("0.0 / 0.0 = %f\n", 0.0/f)   // NaN

	// 溢出
	var a1, b1 int8 = 127, 1
	c1 := a1 + b1
	fmt.Println("\n=== 溢出 === ")
	fmt.Printf("a1 + b1 = %d （二进制：%08b）\n", c1, c1)

	// 实际应用示例
	// 计算用户在xxx网站学习进度的案例

	totalCourses := 51       // 总课程数
	completedCourses := 23   // 已完成课程数
	studyDays := 15          // 学习天数
	dailyStudyMinutes := 120 // 每天学习分钟数

	// 计算学习进度
	progressPercentI := (completedCourses * 100) / totalCourses               // 学习进度百分比，整数比
	progressPercentF := float64(completedCourses*100) / float64(totalCourses) // 学习进度百分比，小数比
	avgCoursesPerDayI := completedCourses / studyDays
	avgCoursesPerDayF := float64(completedCourses) / float64(studyDays)
	totalStudyHoursI := (studyDays * dailyStudyMinutes) / 60
	totalStudyHoursF := float64(studyDays*dailyStudyMinutes) / float64(60)
	avgMinutesPerCourseI := (studyDays * dailyStudyMinutes) / completedCourses
	avgMinutesPerCourseF := float64(studyDays*dailyStudyMinutes) / float64(completedCourses)

	fmt.Println("=== xxx学习统计 === ")
	fmt.Printf("总课程数：%d\n", totalCourses)
	fmt.Printf("已完成：%d\n", completedCourses)
	fmt.Printf("学习进度（整数百分比）：%d%%\n", progressPercentI)
	fmt.Printf("学习进度（小数百分比）：%.2f%%\n", progressPercentF)
	fmt.Printf("平均每天完成：%d 门课程（整数）\n", avgCoursesPerDayI)
	fmt.Printf("平均每天完成：%.2f 门课程（小数）\n", avgCoursesPerDayF)
	fmt.Printf("总学习时长：%d 小时（整数）\n", totalStudyHoursI)
	fmt.Printf("总学习时长：%.2f 小时（小数）\n", totalStudyHoursF)
	fmt.Printf("平均每门课程：%d 分钟（整数）\n", avgMinutesPerCourseI)
	fmt.Printf("平均每门课程：%.2f 分钟（整数）\n", avgMinutesPerCourseF)

	// 预估完成时间
	remainingCourses := totalCourses - completedCourses
	estimateDaysI := int(float64(remainingCourses) / avgCoursesPerDayF)
	estimateDaysF := float64(remainingCourses) / avgCoursesPerDayF

	fmt.Println("\n=== 学习预估 ===")
	fmt.Printf("剩余课程：%d\n", remainingCourses)
	fmt.Printf("预计还需：%d 天完成\n", estimateDaysI)
	fmt.Printf("预计还需：%.2f 天完成\n", estimateDaysF)

	// 关系运算符
	ra, rb, rc := 10, 20, 10
	fmt.Println("\n=== 关系运算符 === ")
	fmt.Printf("ra = %d, rb = %d, rc = %d\n", ra, rb, rc)

	// 相等性比较
	fmt.Printf("ra == rb: %t\n", ra == rb)
	fmt.Printf("ra == rc: %t\n", ra == rc)
	fmt.Printf("ra != rb: %t\n", ra != rb)

	// 大小比较
	fmt.Printf("ra < rb: %t\n", ra < rb)
	fmt.Printf("ra > rb: %t\n", ra > rb)
	fmt.Printf("ra <= rc: %t\n", ra <= rc)
	fmt.Printf("rb >= ra: %t\n", rb >= ra)

	// 字符串比较
	rname1, rname2, rname3 := "张三", "李四", "张三"
	fmt.Printf("\n字符串比较：\n")
	fmt.Printf("'%s' == '%s': %t\n", rname1, rname2, rname1 == rname2)
	fmt.Printf("'%s' == '%s': %t\n", rname1, rname3, rname1 == rname3)
	fmt.Printf("'%s' < '%s': %t\n", rname1, rname2, rname1 < rname2)

	// 浮点数比较的注意事项：因为精度问题可能导致意外结果

	// 浮点数精度问题
	fmt.Println("\n=== 浮点数比较问题 === ")

	rfx, rfy := 0.1, 0.2
	rfa := rfx + rfy
	rfb := 0.3
	fmt.Printf("0.1 + 0.2 = %.17f\n", rfa)
	fmt.Printf("0.3 = %.17f\n", rfb)
	fmt.Printf("0.1 + 0.2 == 0.3: %t\n", rfa == rfb)

	// 正确比较浮点数是否相等：定义一个极小的误差值来比较浮点数是否“足够接近”，而不是直接使用等号
	epsilon := 1e-9                        // 定义一个极小的误差值
	isEqual := math.Abs(rfa-rfb) < epsilon // 比较两数差的绝对值与误差值的大小，如果小于误差值，则认为两数相等
	fmt.Printf("使用epsilon比较：%t\n", isEqual)

	// 实际应用：xxx网站的浮点数计算示例
	userScore1 := 85.6
	userScore2 := 85.6000000001

	fmt.Printf("\n=== xxx评分系统 === ")
	fmt.Printf("用户1得分：%.10f\n", userScore1)
	fmt.Printf("用户2得分：%.10f\n", userScore2)
	fmt.Printf("直接比较相等：%t\n", userScore1 == userScore2)

	// 在评分系统中，通常只关心小数点后几位
	scoreDiff := math.Abs(userScore1 - userScore2)
	isSameSocre := scoreDiff < 0.01
	fmt.Printf("评分系统判定相等：%t\n", isSameSocre)

	// 字符串和字符比较
	fmt.Println("\n=== 字符串比较 === ")

	product1 := "面试鸭"
	product2 := "编程导航"
	product3 := "Face试鸭"

	fmt.Printf("'%s' == '%s': %t\n", product1, product2, product1 == product2)
	fmt.Printf("'%s' < '%s': %t\n", product1, product2, product1 < product2)

	// 字符串长度比较：len()是字节数
	fmt.Printf("'%s' 长度：%d\n", product1, len(product1))
	fmt.Printf("'%s' 长度：%d\n", product3, len(product3))

	// 大小写不敏感比较
	remail1 := "user@mianshiya.com"
	remail2 := "USER@MIANSHIYA.COM"

	fmt.Println("\n=== 邮箱比较 ===")
	fmt.Printf("直接比较：%t\n", remail1 == remail2)
	fmt.Printf("忽略大小写比较：%t\n", strings.EqualFold(remail1, remail2))

	// 字符比较
	fmt.Println("\n=== 字符比较 === ")
	rchar1 := 'A'
	rchar2 := 'a'
	fmt.Printf("'%c' == '%c': %t\n", rchar1, rchar2, rchar1 == rchar2)
	fmt.Printf("'%c' < '%c': %t\n", rchar1, rchar2, rchar1 < rchar2)
	fmt.Printf("'%c' ASCII: %d, '%c' ASCII: %d\n", rchar1, rchar1, rchar2, rchar2)

	// 位运算符
	fmt.Println("\n=== 基本位运算 === ")

	x := 12
	y := 8
	fmt.Printf("x = %d (二进制：%04b)\n", x, x)
	fmt.Printf("y = %d (二进制：%04b)\n", y, y)

	// 按位与
	and := x & y
	fmt.Printf("x & y = %d (二进制：%04b)\n", and, and)

	// 按位或
	or := x | y
	fmt.Printf("x | y = %d (二进制：%04b)\n", or, or)

	// 按位异或
	xor := x ^ y
	fmt.Printf("x ^ y = %d (二进制：%04b)\n", xor, xor)

	// 按位取反
	notX := ^x
	fmt.Printf("^x = %d (二进制：%064b)\n", notX, uint64(notX))

	// 左移
	leftShift := x << 2
	fmt.Printf("x << 2 = %d (二进制：%064b)\n", leftShift, leftShift)

	// 右移
	rightShift := x >> 2
	fmt.Printf("x >> 2 = %d (二进制：%064b)\n", rightShift, rightShift)

	// 位运算的实际应用

	// 位运算的技巧
	fmt.Println("\n2的幂运算：")
	for i := 0; i < 10; i++ {
		power := 1 << i
		fmt.Printf("2^%d = %d\n", i, power)
	}

	// 多重赋值
	text := "你好,世界,GO"
	be, af, fo := stringCut(text, ",")
	fmt.Println(len(text))
	fmt.Printf("分割 '%s': before='%s', after='%s', found=%t\n", text, be, af, fo)
}

func stringCut(s, sep string) (before, after string, found bool) {
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			return s[:i], s[i+len(sep):], true
		}
	}
	return s, "", false
}
