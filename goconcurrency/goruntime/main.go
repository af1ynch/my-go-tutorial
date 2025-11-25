package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(s string, wg *sync.WaitGroup) {

		defer wg.Done()

		for i := 0; i < 2; i++ {
			fmt.Println(s)
			time.Sleep(3 * time.Microsecond)
		}
	}("world", &wg)

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}

	wg.Wait()

}
