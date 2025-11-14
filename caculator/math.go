package caculator

import "math"

// CircleArea 计算圆的面积
func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// Rectangle 计算矩形面积
func Rectangle(height, width float64) float64 {
	return height * width
}

// 私有函数
func square(x float64) float64 {
	return x * x
}

func SquareArea(side float64) float64 {
	return square(side)
}
