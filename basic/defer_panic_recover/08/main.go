package main

import "fmt"

func main() {
	fmt.Println(defer03())
}

func defer03() int {
	defer fmt.Println(1)
	return 2
}

// output:
// 1
// 2

// defer 会在 return 之前执行：
