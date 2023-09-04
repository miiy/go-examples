package main

import "fmt"

func main() {
	case01()
	// case02()
}

func case01() {
	defer fmt.Println(1)
	defer func() {
		panic(2)
	}()
	defer fmt.Println(3)
}

func case02() {
	defer fmt.Println(1)
	defer panic(2)
	defer fmt.Println(3)
}

// case01 and case02 output：
// 3
// 1
// panic: 2
