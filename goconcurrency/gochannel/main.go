package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func send(c chan int, v int) {
	c <- v
	fmt.Println("发送成功", v)
}

func main() {
	// ch1 := make(chan any, 100)
	//
	// ch1 <- 10
	// ch1 <- false
	// ch1 <- []int{1, 2, 3}
	// close(ch1)
	//
	// for item := range ch1 {
	//     fmt.Println(item)
	// }

	ch := make(chan int)
	go send(ch, 10)
	ret := <-ch
	fmt.Println("接收成功", ret)
}
