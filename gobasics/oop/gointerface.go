package oop

import "fmt"

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

type Rectangle struct {
    Height, Width float64
}

func (r Rectangle) Area() float64 {
    return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Height + r.Width)
}

func CaculateShapeInfo(s Shape) {
    fmt.Printf("形状类型：%T\n", s)
    fmt.Printf("面积：%.2f\n", s.Area())
    fmt.Printf("周长：%.2f\n", s.Perimeter())
    fmt.Println("---")
}
