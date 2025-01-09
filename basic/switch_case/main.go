package main

import "fmt"

// 在 case 中添加 fallthrough语句会继续执行紧跟着的下一个case

func main() {
	demo(0)
	fmt.Println("---")
	demo(1)
	fmt.Println("---")
	demo(2)
	fmt.Println("---")
	demo(100)
}

// Output:
//0
//---
//1
//2
//default
//---
//2
//default
//---
//default

func demo(n int) {
	switch n {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("default")
	}
}
