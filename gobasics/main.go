package main

import (
    "fmt"
    "time"
    "unsafe"
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
}

func getSliceMemAddress(c []string) {
    for i := range c {
        fmt.Printf("元素%v的内存地址是%p\n", c[i], &c[i])
    }
}

func getSliceBackingArrayPtr(c []string) *string {
    return unsafe.SliceData(c)
}
