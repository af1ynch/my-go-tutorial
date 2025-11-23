package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	_, err := fmt.Scan(&name, &age, &married)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// fmt.Printf("请输入用户名：")
	// if n, err := fmt.Scanln(&name); err != nil {
	// 	fmt.Printf("%v\n", err)
	// } else {
	// 	fmt.Printf("%v\n", n)
	// }
}
