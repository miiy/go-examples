package main

import (
	"fmt"
	"time"
)

// 预计算参数

func main() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))

	time.Sleep(time.Second)
}

// Output:
// 166ns

// 调用 defer 关键字会立刻拷贝函数中引用的外部参数，所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的，而是在 defer 关键字调用时计算的

// 匿名函数

func main2() {
	startedAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startedAt))
	}()

	time.Sleep(1 * time.Second)
}

// Output:
// 1s

// 虽然调用 defer 关键字时也使用值传递，但是因为拷贝的是函数指针，所以 time.Since(startedAt) 会在 main 函数返回前调用并打印出符合预期的结果。
