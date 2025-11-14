package main

import "fmt"

func main() {
	// switch 的特殊用法

	// 自动 break 和 fallthrough

	grade := 'B'
	fmt.Println("你的成绩评价：")
	switch grade {
	case 'A':
		fmt.Println("优秀")
		fallthrough
	case 'B':
		fmt.Println("良好")
		fallthrough
	case 'C':
		fmt.Println("及格")
		fmt.Println("可以继续学习更高级的内容")
	case 'D':
		fmt.Println("需要补充基础知识")
	default:
		fmt.Println("成绩无效")
	}

	// 多值匹配：一个 case 可以匹配多个值

	month := 13

	switch month {
	case 12, 1, 2:
		fmt.Println("冬天：室内编程的好时节")
		fmt.Println("推荐学习：深入研究算法和数据结构")
	case 3, 4, 5:
		fmt.Println("春天：万物复苏，学习新技术的好时节")
		fmt.Println("推荐学习：新的编程框架和工具")
	case 6, 7, 8:
		fmt.Println("夏季：实习和项目实战的季节")
		fmt.Println("推荐活动：参与开源项目，提升实战能力")
	case 9, 10, 11:
		fmt.Println("秋季：收获的季节，整理和总结")
		fmt.Println("推荐活动：准备面试，完善简历")
	default:
		fmt.Println("一年只有1-12月")
	}

	// type switch
	var data interface{}
	data = 12
	switch v := data.(type) {
	case int:
		fmt.Printf("整数：%d, 平方是：%d\n", v, v*v)
	default:
		fmt.Printf("未知类型：%T，值：%v\n", v, v)
	}
}
