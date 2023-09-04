package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("main defer")
	}()
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
	fmt.Println("test")
}

// Output:
// 打印后
// 打印中
// 打印前
// main defer
// panic: 触发异常

// defer的顺序是后进先出
// 调用 panic 后会立刻停止当前函数的剩余代码，并在当前 Goroutine 中递归执行调用的 defer
