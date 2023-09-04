package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() int {
	defer fmt.Println(1)
	panic(2)
	return 0
}

// Output:
// 1
// panic: 2

// 遇到 panic 会停止执行当前函数后面的代码
