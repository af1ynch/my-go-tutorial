package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	oneHourLater := now.Add(time.Hour)
	fmt.Printf("当前时间是：%v，一个小时后是：%v\n", now.Format(time.DateTime), oneHourLater.Format(time.DateTime))
}
