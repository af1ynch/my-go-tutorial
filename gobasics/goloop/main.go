package main

import "fmt"

func main() {
	// 创建用户技能映射
	userSkills := map[string]int{
		"Go":         8,
		"Python":     7,
		"JavaScript": 6,
		"Docker":     5,
		"Kubernetes": 4,
	}

	fmt.Println("面试鸭用户技能评估：")
	for skill, level := range userSkills {
		fmt.Printf("技能：%-12s 等级：%d/10", skill, level)

		switch {
		case level >= 8:
			fmt.Println(" ⭐⭐⭐ 专家级")
		case level >= 6:
			fmt.Println(" ⭐⭐ 熟练")
		case level >= 4:
			fmt.Println(" ⭐ 基础")
		default:
			fmt.Println(" 待提升")
		}
	}

	// 找出最强和最弱的技能
	var strongestSkill, weakestSkill string
	var maxLevel, minLevel int = 0, 11

	for skill, level := range userSkills {
		if level > maxLevel {
			maxLevel = level
			strongestSkill = skill
		}
		if level < minLevel {
			minLevel = level
			weakestSkill = skill
		}
	}

	fmt.Printf("\n📈 最强技能：%s (等级 %d)\n", strongestSkill, maxLevel)
	fmt.Printf("📉 最弱技能：%s (等级 %d)\n", weakestSkill, minLevel)
	fmt.Println("💡 建议：重点提升薄弱技能，保持优势技能")

	// 在编程导航中搜索特定的课程类型
	fmt.Println("编程导航课程搜索系统：")

	categories := []string{"前端开发", "后端开发", "移动开发", "数据科学"}
	courses := map[string][]string{
		"前端开发": {"HTML基础", "CSS进阶", "JavaScript实战", "React框架"},
		"后端开发": {"Go语言入门", "数据库设计", "API开发", "微服务架构"},
		"移动开发": {"Android开发", "iOS开发", "Flutter跨平台"},
		"数据科学": {"Python数据分析", "机器学习", "深度学习"},
	}

	targetCourse := "Go语言入门"

	// 不使用循环标签时
	fmt.Println("\n=== 未使用循环标签 === ")
	for categoryIndex, category := range categories {
		fmt.Printf("\n搜索分类 %d：%s\n", categoryIndex+1, category)

		courseList := courses[category]
		for courseIndex, course := range courseList {
			fmt.Printf("  检查课程 %d：%s", courseIndex+1, course)

			if course == targetCourse {
				fmt.Printf(" ✅ 找到目标课程！\n")
				fmt.Printf("课程位置：%s -> %s\n", category, course)
				break
			} else {
				fmt.Println(" - 继续搜索")
			}
		}
	}

	fmt.Println("\n=== 使用循环标签 === ")
SearchLoop: // 给外层循环添加标签
	for categoryIndex, category := range categories {
		fmt.Printf("\n搜索分类 %d：%s\n", categoryIndex+1, category)

		courseList := courses[category]
		for courseIndex, course := range courseList {
			fmt.Printf("  检查课程 %d：%s", courseIndex+1, course)

			if course == targetCourse {
				fmt.Printf(" ✅ 找到目标课程！\n")
				fmt.Printf("课程位置：%s -> %s\n", category, course)
				break SearchLoop // 跳出到标签指定的循环
			} else {
				fmt.Println(" - 继续搜索")
			}
		}
	}

	fmt.Println("搜索完成")

	// 面试鸭：多轮面试安排系统
	fmt.Println("面试鸭多轮面试安排：")

	candidates := []string{"张三", "李四", "王五"}
	interviewRounds := []string{"技术初试", "技术复试", "HR面试"}
	timeSlots := []string{"09:00", "10:30", "14:00", "15:30"}

SchedulingLoop: // 主调度循环标签
	for candidateIndex, candidate := range candidates {
		fmt.Printf("\n安排候选人 %s 的面试：\n", candidate)

		scheduledRounds := 0

		// RoundLoop: // 面试轮次循环标签
		for roundIndex, round := range interviewRounds {
			fmt.Printf("  安排 %s：", round)

			// 查找可用时间段
			for slotIndex, slot := range timeSlots {
				// 模拟时间段可用性检查
				isAvailable := (candidateIndex+roundIndex+slotIndex)%3 != 0

				if isAvailable {
					fmt.Printf("时间 %s ✅\n", slot)
					scheduledRounds++
					break // 只跳出时间段循环
				} else {
					fmt.Printf("时间 %s ❌ ", slot)
				}
			}

			// 如果当前轮次没有找到合适时间
			if scheduledRounds <= roundIndex {
				fmt.Printf("\n  ⚠️  %s 无法安排时间，跳过该候选人\n", candidate)
				continue SchedulingLoop // 跳到下一个候选人
			}
		}

		fmt.Printf("  ✅ %s 的所有面试轮次安排完成\n", candidate)
	}

	fmt.Println("\n所有面试安排完成")
}
