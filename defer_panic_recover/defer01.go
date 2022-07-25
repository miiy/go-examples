package main

import "fmt"

func main()  {
	defer01()
}

func defer01()  {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
	fmt.Println("test")
}

// 打印后
// 打印中
// 打印前
// panic: 触发异常
