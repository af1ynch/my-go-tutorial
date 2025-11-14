package oop

import "fmt"

// 接口定义

type Storage interface {
	Save(data string) error
	Load(key string) (string, error)
	Delete(key string) error
}

type FileStorage struct {
	basePath string
}

func (f *FileStorage) Save(data string) error {
	fmt.Printf("将数据保存到文件：%s\n", data)
	return nil
}

func (f *FileStorage) Load(key string) (string, error) {
	fmt.Printf("从文件加载数据：%s\n", key)
	return "文件数据：" + key, nil
}

func (f *FileStorage) Delete(key string) error {
	fmt.Printf("从文件删除数据：%s\n", key)
	return nil
}

// 隐式接口实现

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
