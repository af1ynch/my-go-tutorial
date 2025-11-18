package main

import (
	"fmt"

	"github.com/af1ynch/my-go-tutorial/gobasics/oop/gostruct"
)

func main() {
	var address gostruct.Address

	fmt.Printf("地址信息：%+v\n", address)

	p := new(int)
	p2 := 5
	p3 := &gostruct.User{}
	p4 := new(gostruct.User)

	fmt.Printf("p is %#v\n", p)
	fmt.Printf("p2 is %p\n", &p2)
	fmt.Printf("p3 is %p\n", p3)
	fmt.Printf("p4 is %p\n", p4)

	s := make([]int, 2, 5)
	fmt.Printf("s is %v\n", s)
}
