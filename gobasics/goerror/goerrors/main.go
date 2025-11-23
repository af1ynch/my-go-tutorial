package main

import (
	"errors"
	"fmt"
)

func main() {

	errNotFound := errors.New("not found error")
	fmt.Println(errNotFound)

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v\n", r)
		}
	}()

	s := make([]int, 2, 5)
	fmt.Println(s[3])

}
