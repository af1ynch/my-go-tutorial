/*
数组的定义与初始化
1. 数组的定义与初始化
1.2. 数组的基本定义
2. 数组的访问与遍历
2.1. 基本访问遍历
2.2. 多种遍历方式
3. 多维数组
3.1. 二维数组的创建与使用
3.2. 三维数组的创建与使用
*/
package main

import (
	"fmt"
)

func main() {
	// 定义数组的几种方式
	fmt.Println("=== 定义数组的几种方式 === ")

	// 方式1：声明并初始化
	var scores [5]int = [5]int{95, 98, 100, 95, 70}
	fmt.Printf("用户成绩：%v\n", scores)

	// 方式2：类型推断初始化
	languages := [3]string{"Go", "Python", "JavaScript"}
	fmt.Printf("热门编程语言：%v\n", languages)

	// 方式3：编译器自动推断长度
	frameworks := [...]string{"Gin", "Echo", "FastAPI"}
	fmt.Printf("Python和Golang常用Web框架：%v (长度：%d)\n", frameworks, len(frameworks))

	// 方式4：指定索引初始化
	weekdays := [7]string{0: "星期天", 1: "星期一", 6: "星期六"}
	fmt.Printf("工作日安排：%v\n", weekdays)

	// 方式5：零值初始化
	var userLevels [10]int
	fmt.Printf("用户等级列表：%v\n", userLevels)

	// 实际应用：简历模板管理
	var templateCategories [6]string
	templateCategories[0] = "技术类"
	templateCategories[1] = "设计类"
	templateCategories[2] = "销售类"
	templateCategories[3] = "管理类"
	templateCategories[4] = "教育类"
	templateCategories[5] = "其他类"

	fmt.Println("\n=== 简历模板分类 === ")
	for i, category := range templateCategories {
		fmt.Printf("%d. %s\n", i+1, category)
	}

	// 每个分类的模板数量统计
	templateCounts := [6]int{25, 18, 12, 15, 10, 8}

	fmt.Println("=== 各分类模板数量 === ")
	totalTemplates := 0
	for i, count := range templateCounts {
		fmt.Printf("%s模板数量：%d\n", templateCategories[i], count)
		totalTemplates += count
	}
	fmt.Printf("总计：%d个模板\n", totalTemplates)

	// 找出最受欢迎的模板分类
	maxCount := 0
	popularCategory := ""
	for i, count := range templateCounts {
		if count > maxCount {
			maxCount = count
			popularCategory = templateCategories[i]
		}
	}
	fmt.Printf("最受欢迎的模板分类是：%s\n", popularCategory)

	// 数组的访问与遍历

	// 基本访问操作

	// 课程评分数组
	courseRatings := [5]float64{4.8, 4.9, 4.7, 4.6, 4.9}
	courseNames := [5]string{"Go基础", "算法入门", "项目实战", "面试准备", "职业规划"}

	fmt.Println("\n=== 课程评分 === ")
	// 基本访问：通过索引
	fmt.Printf("第一门课程：%s，评分：%.1f\n", courseNames[0], courseRatings[0])
	fmt.Printf("最后一门课程：%s，评分：%.1f\n", courseNames[len(courseNames)-1], courseRatings[len(courseRatings)-1])

	// 数组的长度和容量相等
	users := [5]string{"张三", "李四", "王五"}
	fmt.Printf("课程总数：%d\n", len(users))
	fmt.Printf("数组容量：%d\n", cap(users))

	// 多种遍历方式
	skills := [6]string{"Go", "Docker", "K8s", "MySQL", "Redis", "Git"}
	skillLevels := [6]int{9, 7, 6, 8, 7, 9}

	fmt.Println("\n=== 剪切助手开发团队技能评估 === ")

	// 方式一：传统 for 循环遍历
	fmt.Println("\n方式一：传统 for 循环遍历")
	length := len(skillLevels)
	for i := 0; i < length; i++ {
		fmt.Printf("技能 %d: %s (熟练度: %d/10)\n", i+1, skills[i], skillLevels[i])
	}

	// 方式二：range 遍历（获取索引和值）
	fmt.Println("\n方式二：range 遍历")
	for index, skill := range skills {
		level := skillLevels[index]
		status := ""
		switch {
		case level >= 8:
			status = "精通 ⭐️⭐️⭐️"
		case level >= 6:
			status = "熟练 ⭐️⭐️"
		default:
			status = "基础 ⭐️"
		}
		fmt.Printf("%s: %s\n", skill, status)
	}

	// 方式三：只遍历值，忽略索引
	fmt.Println("\n方式三：只获取值")
	for _, skill := range skills {
		fmt.Printf("- %s\n", skill)
	}
	fmt.Println()

	// 方式四：只遍历索引，忽略值
	fmt.Println("\n方式四：只获取索引")
	for index := range skills {
		fmt.Printf("位置 %d ", index)
	}
	fmt.Println()

	// 数组比较
	userScores1 := [3]int{100, 100, 95}
	userScores2 := [3]int{100, 100, 95}
	fmt.Printf("user1分数 == user2分数：%t\n", userScores1 == userScores2)

	// 多维数组

	// 二维数组的创建与使用

	// 题库难度分布矩阵
	// 行代表知识点，列代表难度等级
	questionMatrix := [4][3]int{
		{15, 8, 3},
		{12, 10, 5},
		{20, 6, 2},
		{8, 12, 8}}

	categories := [4]string{"数据结构", "算法设计", "基础语法", "系统设计"}
	difficulties := [3]string{"简单", "中等", "苦难"}

	fmt.Println("\n=== 题库分布统计 === ")
	fmt.Printf("%-12s", "知识点\\难度")
	for _, diff := range difficulties {
		fmt.Printf("%-8s", diff)
	}
	fmt.Println("总计")

	totalByCategory := [4]int{}
	totalByDifficulty := [3]int{}
	grandTotal := 0

	for i, category := range categories {
		fmt.Printf("%-12s", category)
		categoryTotal := 0

		for j, count := range questionMatrix[i] {
			fmt.Printf("%-8d", count)
			categoryTotal += count
			totalByDifficulty[j] += count
			grandTotal += count
		}

		totalByCategory[i] = categoryTotal
		fmt.Printf("%d\n", categoryTotal)
	}
}
