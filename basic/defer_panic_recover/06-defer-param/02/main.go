package main

import "fmt"

func main() {
	test()
}

func test() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// Output:
// 0

// 传入 defer的函数会被先计算
