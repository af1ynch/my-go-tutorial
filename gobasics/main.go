package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unsafe"

	"github.com/af1ynch/my-go-tutorial/gobasics/oop/gointerface"
	"github.com/af1ynch/my-go-tutorial/gobasics/oop/gostruct"
)

// 定义同类型多全局变量
var (
	name10 = "Jack"
	name11 = "Chan"
)

func main() {
	// Hello, World
	fmt.Println("Hello World")

	// 变量
	// 方式一：var var_name type = value，声明时可以不初始化（即没有值），变量默认为零值
	var bool1 bool   // 零值为：false
	var num1 int     // 0
	var num2 float32 // 0
	var byte1 byte   // 0
	var rune1 rune   // 0
	var str1 string  // ""
	var ptr1 *int
	var slice1 []int
	var map1 map[string]string
	fmt.Println(bool1, num1, num2, byte1, rune1)
	fmt.Println(str1)
	fmt.Println(ptr1)
	fmt.Println(slice1)
	fmt.Println(map1)
	fmt.Printf("指针的零值: %v, 是否为 nil: %t\n", ptr1, ptr1 == nil)
	fmt.Printf("切片的零值: %#v, 是否为 nil: %t\n", slice1, slice1 == nil)
	fmt.Printf("map 的零值: %#v, 是否为 nil: %t\n", map1, map1 == nil)

	// 方式二：初始化声明变量时可以省略类型，编译器会根据值自行判断类型
	var str2 = "hello,world"
	fmt.Println(str2)

	// 方式三：短赋值语句声明局部变量，var_name := value 等价于 var var_name string = "hello,world"
	str3 := "hello,world"
	fmt.Println(str3)

	// 类型相同的多变量声明
	// 如果是局部变量
	// var name1, name2, name3 type name1. name2, name3 = value1, value2, value3
	// var name1, name2, name3 = value1, value2, value3
	// name1, name2, name3 := value2, value2, value3
	var name1, name2, name3 string
	name1, name2, name3 = "hello", "world", "!"
	fmt.Println(name1, name2, name3)

	var name4, name5, name6 = "hello", "world", "!"
	fmt.Println(name4, name5, name6)

	name7, name8, name9 := "hello", "world", "!"
	fmt.Println(name7, name8, name9)

	// 如果是全局变量
	// 可以用因式分解关键字的方法
	/*
	   	   var (
	               name1 type1
	               name2 type2
	   	   )
	*/
	fmt.Println(name10, name11)

	// 常量
	const pi = 3.14
	fmt.Println("pi:", pi)

	// 枚举
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("Sunday:", Sunday)

	// 常量表达式
	// const (
	//     a = "abc"
	//     b = len(a)
	//     c = unsafe.Sizeof(a)
	// )
	// fmt.Printf("a: %s, b: %d, c: %d\n", a, b, c)

	// iota
	const (
		a = iota
		b
		c
		d = "haha"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	// 浮点数
	var myfloat float32 = 100000187
	fmt.Println(myfloat)
	fmt.Println(myfloat + 1)

	// byte 类型
	var mybyte byte = 'a'
	var myint uint8 = 97
	fmt.Printf("%c\n", mybyte)
	fmt.Printf("%c\n", myint)

	// rune 类型
	var myrune = '中'
	fmt.Printf("%c\n", myrune)

	// 字符串
	var mystr01 = "hello"
	var mystr02 = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s\n", mystr02)

	mystr03 := `这是
一首
    简单`
	fmt.Println(mystr03)

	// 格式化输出
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", data)
	fmt.Printf("%%#v：%#v\n", data)

	// 条件语句
	/*
	   if condition {
	   } else if condition {
	   } else {
	   }
	*/
	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// 循环语句
	/*
	   	   for init expresison; condition expression; update expression {
	       }
	*/
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}

	// 迭代：range
	arr02 := [...]int{1, 2, 3, 4, 5}
	ptr2 := &arr02
	for _, v := range *ptr2 {
		fmt.Println(v)
	}

	// switch
	num := 2
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("outer")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// 数组

	// 定义
	// 一维数组：var name [size]type
	var course1 [3]string
	// 二维数组：var name [size][size]type
	var courseInfo [3][4]string

	// 初始化
	// 一维数组
	// 方法一：数组字面值语法初始化
	course1 = [3]string{"Go", "Python", "Java"}
	// 方法二：使用索引初始化
	course2 := [3]string{2: "Java"}
	// 方法三：使用 ... 初始化，...表示根据初始化值的数量自动推断数组的长度
	course3 := [...]string{"Go", "Python", "Java"}
	fmt.Println(course1)
	fmt.Println(course2)
	fmt.Println(course3)

	// 多维数组
	// 方法一：数组字面值语法初始化
	courseInfo = [3][4]string{
		{"Go", "Golang", "Go语言", "Go 语言"},
		{"Python", "Python3", "Python 3", "Python 3 语言"},
		{"Java", "JavaSE", "Java 标准 Edition", "Java 标准版"},
	}
	fmt.Println(courseInfo)

	// 方法二：使用键控元素初始化，非常灵活允许我们在初始化时指定行索引和/或列索引
	courseInfo3 := [3][4]string{
		0: {"Go", "Golang", "Go语言", "Go 语言"},
		1: {"Python", "Python3", "Python 3", "Python 3 语言"},
		2: {"Java", "JavaSE", "Java 标准 Edition", "Java 标准版"},
	}
	fmt.Println(courseInfo3)
	// 方法四：逐个元素赋值
	var courseInfo4 [3][4]string
	courseInfo4[0][0] = "Go"
	courseInfo4[0][1] = "Golang"
	courseInfo4[0][2] = "Go语言"
	courseInfo4[0][3] = "Go 语言"

	months := [...]string{1: "January", 12: "December"}
	fmt.Println(months)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	// 切片

	// 定义
	// var name []type

	// 初始化
	// 方法一：字面值
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	// 方法二：make
	slice3 := make([]int, 3, 5)
	fmt.Printf("%p, %v, len: %d, cap: %d\n", &slice3, slice3, len(slice3), cap(slice3))

	// 方法三：从数组或切片直接创建
	slice4 := course1[0:2]
	fmt.Println(slice4, len(slice4), cap(slice4))

	// 方法四：append
	var slice5 []string
	slice5 = append(slice5, "Go")
	fmt.Println(slice5, len(slice5), cap(slice5))

	// 常用操作

	// 增加元素
	// 场景一：增加单个元素
	var course4 []string
	course4 = append(course4, "Go")
	fmt.Println(course4, len(course4), cap(course4))

	// 场景二：增加多个元素
	course4 = append(course4, "Python", "Java")
	fmt.Println(course4, len(course4), cap(course4))

	// 删除元素
	// 元素删除，容量不变
	course5 := []string{"Go", "Python", "Java", "grpc", "gin"}
	fmt.Printf("%p\n", &course5)
	fmt.Printf("course5当前底层数组的指针是：%v\n", getSliceBackingArrayPtr(course5))
	getSliceMemAddress(course5)
	course5 = append(course5[0:2], course5[3:]...)
	fmt.Printf("%p\n", &course5)
	fmt.Printf("course5删除元素后底层数组的指针是：%p\n", getSliceBackingArrayPtr(course5))
	getSliceMemAddress(course5)
	fmt.Println(course5, len(course5), cap(course5))
	course5 = append(course5, "mysql")
	fmt.Printf("%p\n", &course5)
	fmt.Printf("course5添加元素后底层数组的指针是：%p\n", getSliceBackingArrayPtr(course5))
	getSliceMemAddress(course5)
	fmt.Println(course5, len(course5), cap(course5))

	// 拷贝

	// 浅拷贝：切片改动会影响到原切片
	course6 := []string{"Golang", "Python", "C"}
	fmt.Println(course6, len(course6), cap(course6))
	fmt.Println()

	course7 := course6[:]
	fmt.Printf("%p\n", &course6)
	fmt.Printf("%p\n", &course7)
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course6))
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course7))
	fmt.Println(course6, len(course6), cap(course6))
	fmt.Println(course7, len(course7), cap(course7))
	course7[0] = "Java"
	fmt.Printf("%p\n", &course6)
	fmt.Printf("%p\n", &course7)
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course6))
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course7))
	fmt.Println(course6, len(course6), cap(course6))
	fmt.Println(course7, len(course7), cap(course7))
	course7 = append(course7, "mysql")
	fmt.Printf("%p\n", &course6)
	fmt.Printf("%p\n", &course7)
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course6))
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course7))
	fmt.Println(course6, len(course6), cap(course6))
	fmt.Println(course7, len(course7), cap(course7))

	// 深拷贝：切片改动不会影响到原切片
	course8 := make([]string, len(course6), cap(course6))
	copy(course8, course6)
	fmt.Printf("%p\n", &course6)
	fmt.Printf("%p\n", &course8)
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course6))
	fmt.Printf("%p\n", getSliceBackingArrayPtr(course8))

	// make
	courses9 := make([]int, 2, 5)
	courses10 := make([]string, 2, 5)
	courses11 := make([]bool, 2, 5)
	fmt.Println(courses9, len(courses9), cap(courses9))
	fmt.Println(courses10, len(courses10), cap(courses10))
	fmt.Println(courses11, len(courses11), cap(courses11))

	// new
	course12 := new([]string)
	fmt.Println(course12, len(*course12), cap(*course12))

	// map

	// 初始化
	// 方法一：字面值初始化
	var courseMap1 = map[string]string{
		"mysql":  "mysql入门",
		"python": "python入门",
		"go":     "go入门",
	}
	fmt.Printf("courseMap1的内存地址是%p, %v\n", &courseMap1, courseMap1)

	// 方法二：空字面值初始化
	var courseMap2 = map[string]string{}
	fmt.Printf("courseMap2的内存地址是%p, %v\n", &courseMap2, courseMap2)
	courseMap2["java"] = "java入门"

	// 方法三：make函数初始化，第二个参数是容量提示（可选），有助于性能优化
	var courseMap3 = make(map[string]string, 3)
	courseMap3["go"] = "Golang入门"
	fmt.Printf("courseMap3的内存地址是%p, %v\n", &courseMap3, courseMap3)

	// 注意：
	// var courseMap4 = map[string]string
	// courseMap4["mysql"] = "mysql入门"

	// 判断KV是否存在
	value, ok := courseMap3["go"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("go课程不存在")
	}

	if _, ok := courseMap3["go"]; ok {
		fmt.Println("go课程存在")
	} else {
		fmt.Println("go课程不存在")
	}

	// 删除元素
	delete(courseMap3, "go")

	// 指针

	// 声明一个普通变量
	userName := "周星驰"

	// 声明一个指向普通变量的指针
	userPointer := &userName

	fmt.Printf("用户名：%s\n", userName)
	fmt.Printf("用户名地址：%p\n", &userName)
	fmt.Printf("指针存储的地址：%p\n", userPointer)
	fmt.Printf("通过指针访问的值：%s\n", *userPointer)

	var ptr *int

	fmt.Printf("ptr的内存地址是%p, %v\n", &ptr, ptr)
	// fmt.Printf("ptr指向的值是%v\n", *ptr)

	if ptr != nil {
		fmt.Printf("ptr指向的值是%v\n", *ptr)
	} else {
		fmt.Println("指针为空，无法解引用")
	}

	// 指针初始化

	// 先声明变量，再取指针
	resumeTitle := "软件工程师"
	resumePointer := &resumeTitle

	fmt.Printf("resumeTitle的内存地址是%p, %v\n", &resumeTitle, resumeTitle)
	fmt.Printf("resumePointer的内存地址是%p, %v\n", &resumePointer, resumePointer)
	fmt.Printf("通过指针访问的值：%s\n", *resumePointer)

	// 使用 new 函数分配内存
	salary := new(float64)
	*salary = 5000.00
	fmt.Printf("salary的内存地址是%p, %v\n", &salary, *salary)

	// 指针的指针
	originText := "Hello, 剪切助手"
	textPointer := &originText
	doublePtr := &textPointer

	fmt.Printf("doublePtr的内存地址是%p, %v\n", &doublePtr, doublePtr)
	fmt.Printf("通过指针访问的值：%s\n", **doublePtr)

	// 函数
	sum1, _ := add(1, 2)
	fmt.Println(sum1)

	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))

	test1 := 1
	test2 := 2
	test3 := add2(&test1, &test2)
	fmt.Println(test1, test3)

	// 结构体与方法

	var user1 gostruct.User

	// 初始化

	// 字段逐一赋值
	user1.ID = 1001
	user1.Name = "Alice"
	user1.Email = "alice@example.com"
	user1.Age = 30
	user1.IsActive = true

	// 打印用户信息
	fmt.Printf("用户信息如下：%+v\n", user1)
	fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, IsActive: %v\n",
		user1.ID, user1.Name, user1.Email, user1.Age, user1.IsActive)

	// 使用字面量初始化

	// 按字段顺序初始化，顺序必须一致
	// user2 := gostruct.User{ "Bob", "bob@example.com", 25, gostruct.Address{"123 Main St", "Anytown", "CA", "12345"}, true}

	// 打印用户信息
	// fmt.Printf("用户信息如下：%+v\n", user2)
	// fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, IsActive: %v\n",
	// 	user2.ID, user2.Name, user2.Email, user2.Age, user2.IsActive)

	// 按字段名初始化，顺序可以不一致
	user3 := gostruct.User{
		Name:     "Bob",
		Email:    "bob@example.com",
		Age:      25,
		Address:  gostruct.Address{"123 Main St", "Anytown", "CA", "12345"},
		IsActive: true,
	}

	// 打印用户信息
	fmt.Printf("用户信息如下：%+v\n", user3)
	fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, Address: %+v, IsActive: %v\n",
		user3.ID, user3.Name, user3.Email, user3.Age, user3.Address.State+" "+user3.Address.City, user3.IsActive)

	// 部分字段初始化，未初始化的字段将使用零值
	user4 := gostruct.User{
		Name:     "Charlie",
		Email:    "charlie@example.com",
		IsActive: false,
	}

	// 打印用户信息
	fmt.Printf("用户信息如下：%+v\n", user4)
	fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, IsActive: %v\n",
		user4.ID, user4.Name, user4.Email, user4.Age, user4.IsActive)

	// 使用 new 函数
	user5 := new(gostruct.User)
	user5.Name = "David"
	user5.Email = "david@example.com"
	user5.Age = 28
	user5.IsActive = true

	// 打印用户信息
	fmt.Printf("用户信息如下：%+v\n", *user5)
	fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, IsActive: %v\n",
		user5.ID, user5.Name, user5.Email, user5.Age, user5.IsActive)

	// 使用构造函数初始化
	user6 := gostruct.NewUser("Eve", "eve@example.com")

	// 打印用户信息
	fmt.Printf("用户信息如下：%+v\n", *user6)
	fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, IsActive: %v\n",
		user6.ID, user6.Name, user6.Email, user6.Age, user6.IsActive)

	acc := gostruct.NewAccount("123456", "Alice")
	fmt.Printf("acc的余额是%v\n", acc.GetBalance())

	newBalance := acc.Deposit(100)
	fmt.Printf("存款后返回的余额是：%.2f\n", newBalance)
	fmt.Printf("存款后账户余额是：%.2f\n", acc.GetBalance())

	doc1 := gostruct.CreateDocument("Go语言基础")
	doc1.AppendContent("Go语言是一种静态类型、编译型语言，由Google开发。")
	fmt.Printf("文档标题是：%s\n", doc1.GetTitle())
	fmt.Printf("文档大小是：%d\n", doc1.GetSize())

	// 序列化为JSON
	course13 := gostruct.Course{
		ID:          1001,
		Title:       "Go语言基础",
		Description: "介绍Go语言的基础语法、数据类型、控制流程等。",
		Price:       99.99,
		IsPublished: false,
		Tags:        []string{"Go", "语法", "基础"},
		Author:      "张三",
	}
	fmt.Printf("序列化前的课程信息：\n%#v\n", course13)
	jsonData, err := json.Marshal(course13)
	if err != nil {
		fmt.Println("JSON序列化错误:", err)
		return
	}
	fmt.Printf("课程信息如下：\n%s\n，数据类型是%T\n", jsonData, jsonData)

	// 接口
	circle1 := &gointerface.Circle{Radius: 2}
	gointerface.CaculateShapeInfo(circle1)

	rectangle1 := &gointerface.Rectangle{
		Height: 3,
		Width:  2,
	}
	gointerface.CaculateShapeInfo(rectangle1)
}

func getSliceMemAddress(c []string) {
	for i := range c {
		fmt.Printf("元素%v的内存地址是%p\n", c[i], &c[i])
	}
}

func getSliceBackingArrayPtr(c []string) *string {
	return unsafe.SliceData(c)
}

func add(a, b int) (sum int, err error) {
	sum = a + b
	return
}

func swap(a, b int) (int, int) {
	return b, a
}

func divide(a, b float64) (res float64, err error) {
	if b == 0.0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

func add1(desc string, items ...int) (sum int, err error) {
	for _, value := range items {
		sum += value
	}
	return sum, err
}

func calculate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 引用传递
func add2(a, b *int) int {
	*a += 10
	return *a + *b
}
