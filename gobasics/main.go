package main

import (
    "fmt"
    "time"
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
    myarr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(myarr)

    myarr01 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(myarr01)

    // 切片
    arr01 := [5]int{1, 2, 3, 4, 5}
    fmt.Printf("arr01 的长度为：%d, 容量为：%d\n", len(arr01), cap(arr01))
    sli01 := arr01[1:3]
    fmt.Printf("sli01 的长度为：%d, 容量为：%d\n", len(sli01), cap(sli01))
}
