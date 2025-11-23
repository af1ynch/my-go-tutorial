package main

import (
	"fmt"
	"reflect"
)

type Myint = int
type Myint2 Myint

func main() {
	logs := []any{"User", 123, "login successfully"}
	fmt.Println(logs...)

	var s []int = nil
	fmt.Printf("%v\n", s)

	var b byte = 'c'
	var u8 uint8 = 1
	fmt.Printf("b的类型%T == u8的类型%T\n", b, u8)

	var a Myint = 1
	c := 2
	var d Myint2 = 2
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d) == reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(d))

	bb()
}

func f(a ...int) {
	fmt.Printf("%v", a)
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func bb() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
