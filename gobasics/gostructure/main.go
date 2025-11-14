package main

import (
	"fmt"
	"math"
)

const PI = 3.141592654

var appName = "第一个 Go 程序"

func init() {
	fmt.Println("init 函数被调用")
}

func main() {
	radius := 5.0

	area := PI * math.Pow(radius, 2)

	fmt.Printf("%s\n", appName)
	fmt.Printf("半径为：%.1f 的圆的面积是：%.2f\n", radius, area)

}
